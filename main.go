package main

import (
	"fmt"
	"net/http"

	"github.com/SlavaUtesinov/store/handlers"
)

func main() {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handlers.CreateHandler(),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error has happened during server startup: %v", err)
	}
}
