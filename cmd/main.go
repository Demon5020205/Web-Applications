package main

import (
	"log"

	"example.com/mod/pkg/handler"
	"example.com/mod/pkg/repository"
	"example.com/mod/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error  running http server: %s", err.Error)
	}
}