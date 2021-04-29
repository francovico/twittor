package main

import (
	"log"

	"github.com/francovico/twittor/bd"
	"github.com/francovico/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
