package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"bestWebApp/internal/server"
)

func main() {
	config := &server.Config{
		Port:         "8088",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server := server.NewServer(config)
	if err := server.Start(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}
