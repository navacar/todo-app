package main

import (
	"log"
	todo "rest-api-todo"
	"rest-api-todo/pkg/handler"
)

func main() {
	handler := handler.New()
	srv := todo.New()
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Error starting server: %w", err)
	}
}
