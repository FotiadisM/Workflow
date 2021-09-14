package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/FotiadisM/workflow-server/internal/auth"
	"github.com/FotiadisM/workflow-server/internal/conversations"
	"github.com/FotiadisM/workflow-server/internal/jobs"
	"github.com/FotiadisM/workflow-server/internal/posts"
	"github.com/FotiadisM/workflow-server/internal/repository"
	"github.com/FotiadisM/workflow-server/internal/user"
	"github.com/FotiadisM/workflow-server/pkg/middleware"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Session struct {
	UserID string
}

var (
	httpPort string = "8080"
	dbURL    string = "postgresql://root@localhost:26257/workflow?sslmode=disable"
)

func init() {
	if p := os.Getenv("PORT"); p != "" {
		httpPort = p

	}

	if url := os.Getenv("DB_URL"); url != "" {
		dbURL = url
	}
}

// interruptHandler handles graceful shutdown
func interruptHandler(errc chan<- error, httpServer *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	errc <- httpServer.Shutdown(context.Background())
}

func main() {

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)

	ctx := context.Background()
	repo, err := repository.NewRepository(ctx, dbURL)
	if err != nil {
		panic(fmt.Errorf("Failed at creating Repository: %w", err))
	}

	fp := repository.FilesPath
	if err = os.MkdirAll(fp, 0777); err != nil {
		panic(fmt.Errorf("failed to create folder for files %w", err))
	}

	options := []httptransport.ServerOption{
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	var clients = make(map[*websocket.Conn]Session)
	var broadcast = make(chan conversations.BroadCastMessage)
	r := mux.NewRouter()
	{
		autSvc := auth.NewService(repo)
		authEnds := auth.NewEndpoints(autSvc)
		authEnds.WrapAllExcept(middleware.Logger(logger))
		auth.NewHTTPRouter(authEnds, r.PathPrefix("/auth").Subrouter(), options...)

		userSvc := user.NewService(repo)
		userEnds := user.NewEndpoints(userSvc)
		user.NewHTTPHandler(userEnds, r.PathPrefix("/users").Subrouter(), options...)

		postsSvc := posts.NewService(repo)
		postsEnds := posts.NewEndpoints(postsSvc)
		posts.NewHTTPRouter(postsEnds, r.PathPrefix("/posts").Subrouter(), options...)

		convSvc := conversations.NewService(repo, broadcast)
		convEnds := conversations.NewEndpoints(convSvc)
		conversations.NewHTTPRouter(convEnds, r.PathPrefix("/conversations").Subrouter(), options...)

		jobsSvc := jobs.NewService(repo)
		jobEnds := jobs.NewEndpoints(jobsSvc)
		jobs.NewHTTPRouter(jobEnds, r.PathPrefix("/jobs").Subrouter(), options...)

		r.HandleFunc("/static/{name}", func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			http.ServeFile(rw, r, filepath.Join(repository.FilesPath, vars["name"]))
		})

		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		r.HandleFunc("/ws/{userID}", func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userID := vars["userID"]
			conn, err := upgrader.Upgrade(rw, r, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			clients[conn] = Session{
				UserID: userID,
			}

			for {
				var msg conversations.Message
				err = conn.ReadJSON(&msg)
				if err != nil {
					fmt.Println(err)
					delete(clients, conn)
					return
				}
				// TODO: store message to db
				// broadcast <- msg
			}
		})
	}

	co := handlers.AllowedOrigins([]string{"http://localhost:3000", "*"})
	ch := handlers.AllowedHeaders([]string{"Content-Type", "*"})
	cm := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "*"})

	httpServer := &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: handlers.CORS(co, ch, cm)(r),
	}

	errc := make(chan error)

	go func() {
		logger.Log("listening", httpPort)
		errc <- httpServer.ListenAndServe()
	}()

	go func() {
		for {
			msg := <-broadcast

			for conn, session := range clients {
				if msg.Receiver == session.UserID {
					if err := conn.WriteJSON(msg); err != nil {
						fmt.Println(err)
						conn.Close()
						delete(clients, conn)
					}
				}
			}
		}
	}()

	go interruptHandler(errc, httpServer)

	logger.Log("exit", <-errc)
}
