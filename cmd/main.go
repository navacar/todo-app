package main

import (
	"log"
	todo "rest-api-todo"
	"rest-api-todo/pkg/handler"
	"rest-api-todo/pkg/repository"
	"rest-api-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.New(services)

	srv := todo.New()
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting server: %w", err)
	}
}
