package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bestWebApp/internal/handler"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Server struct {
	*http.Server
	config *Config
}

func NewServer(config *Config) *Server {
	h := handler.New()

	srv := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      h,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return &Server{
		Server: srv,
		config: config,
	}
}

func (s *Server) Start() error {
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is starting on port %s", s.config.Port)
		serverErrors <- s.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("error starting server: %w", err)

	case sig := <-shutdown:
		log.Printf("Start shutdown... Signal: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			s.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
