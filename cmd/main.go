package main

import (
	"github.com/Ownax-vit/simpeRestGo"
	"github.com/Ownax-vit/simpeRestGo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
