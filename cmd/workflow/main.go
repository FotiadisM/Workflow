package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/FotiadisM/workflow-server/internal/jobs"
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

	s := jobs.NewService(nil)
	e := jobs.NewEndpoints(s)
	jobs.NewHTTPRouter(e, r.PathPrefix("/jobs").Subrouter())

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
