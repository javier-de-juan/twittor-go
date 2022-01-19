package main

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/handlers"
	"log"
)

func main() {
	if ! bd.IsConnected() {
		log.Fatal("Database is not connected")
		return
	}

	handlers.Handle()
}
