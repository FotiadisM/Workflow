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
	"github.com/gorilla/mux"
)

var (
	httpPort string = "8080"
)

func init() {
	if p := os.Getenv("PORT"); p != "" {
		httpPort = p
	}
}

func main() {
	errc := make(chan error)

	var httpServer *http.Server

	r := mux.NewRouter()

	repo := repository.NewRepository()

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
	jobs.NewHTTPRouter(jobEnds, r.PathPrefix("/jobs").Subrouter())

	httpServer = &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: r,
	}

	go func() {
		log.Println("server listening on port:", httpPort)
		errc <- httpServer.ListenAndServe()
	}()

	err := interruptHandler(errc, httpServer)
	if err != nil {
		log.Println("error shuting down:", err)
	}

	log.Println("exit", <-errc)
}

// interruptHandler handles graceful shutdown
func interruptHandler(errc chan<- error, httpServer *http.Server) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	return httpServer.Shutdown(context.Background())
}
