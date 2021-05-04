package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/FotiadisM/workflow-server/internal/auth"
	"github.com/FotiadisM/workflow-server/internal/conversations"
	"github.com/FotiadisM/workflow-server/internal/jobs"
	"github.com/FotiadisM/workflow-server/internal/posts"
	"github.com/FotiadisM/workflow-server/internal/repository"
	"github.com/FotiadisM/workflow-server/internal/user"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

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

func main() {
	ctx := context.Background()
	repo, err := repository.NewRepository(ctx, dbURL)
	if err != nil {
		log.Println("error connecting to databse:", err)
		return
	}

	options := []httptransport.ServerOption{
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	r := mux.NewRouter()
	{
		autSvc := auth.NewService(repo)
		authEnds := auth.NewEndpoints(autSvc)
		auth.NewHTTPRouter(authEnds, r.PathPrefix("/auth").Subrouter())

		userSvc := user.NewService(nil)
		userEnds := user.NewEndpoints(userSvc)
		user.NewHTTPHandler(userEnds, r.PathPrefix("/users").Subrouter())

		postsSvc := posts.NewService(nil)
		postsEnds := posts.NewEndpoints(postsSvc)
		posts.NewHTTPRouter(postsEnds, r.PathPrefix("/posts").Subrouter())

		convSvc := conversations.NewService(nil)
		convEnds := conversations.NewEndpoints(convSvc)
		conversations.NewHTTPRouter(convEnds, r.PathPrefix("/conversations").Subrouter())

		jobsSvc := jobs.NewService(nil)
		jobEnds := jobs.NewEndpoints(jobsSvc)
		jobs.NewHTTPRouter(jobEnds, r.PathPrefix("/jobs").Subrouter(), options...)
	}

	httpServer := &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: r,
	}

	errc := make(chan error)

	go func() {
		log.Println("server listening on port:", httpPort)
		errc <- httpServer.ListenAndServe()
	}()

	go interruptHandler(errc, httpServer)

	log.Println("exit", <-errc)
}

// interruptHandler handles graceful shutdown
func interruptHandler(errc chan<- error, httpServer *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	errc <- httpServer.Shutdown(context.Background())
}
