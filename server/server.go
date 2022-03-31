package server

import (
	"fmt"
	"net/http"

	"github.com/SlavaUtesinov/store/handlers"
)

func Run(port int) error {
	addr := fmt.Sprintf("localhost:%v", port)
	server := http.Server{
		Addr:    addr,
		Handler: handlers.CreateHandler(),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error has happened during server startup: %v", err)
		return err
	}

	return nil
}
