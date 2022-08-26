package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server is a base server configuration.
type server struct {
	*http.Server
}

// newServer initialize a new server with configuration.
func newServer(listening string, mux http.Handler) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *server) Start() {
	log.Println("starting API cmd")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("could not listen on" + srv.Addr + "due to " + err.Error())
		}
	}()
	log.Println("cmd is ready to handle requests ", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("cmd is shutting down ", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("could not gracefully shutdown the cmd ", err.Error())
	}
	log.Println("cmd stopped")
}
