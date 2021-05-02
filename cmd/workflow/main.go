package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/FotiadisM/workflow-server/internal/user"
)

var (
	httpPort string = ":8080"
)

func init() {
	if p := os.Getenv("PORT"); p != "" {
		httpPort = ":" + p
	}
}

func main() {
	errc := make(chan error)

	var httpServer *http.Server

	s := user.NewService()
	e := user.NewEndpoints(s)
	h := user.NewHTTPHandler(e)

	httpServer = &http.Server{
		Addr:    httpPort,
		Handler: h,
	}

	go func() {
		log.Println("server listening on port:", httpPort)

		errc <- httpServer.ListenAndServe()
	}()

	interruptHandler(errc, httpServer)

	log.Println("exit", <-errc)
}

// interruptHandler handles graceful shutdown
func interruptHandler(errc chan<- error, httpServer *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	httpServer.Shutdown(context.Background())
}
