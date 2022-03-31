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
	if error := server.ListenAndServe(); error != nil {
		fmt.Printf("Error has happened: %v", error)
	}
}
