package main

import (
	"fmt"
	"log"

	"github.com/fake_twitter/src/bd"
	"github.com/fake_twitter/src/handlers"
)

func main() {
	fmt.Println("Hola")
	if bd.CheckConnection() == false {
		log.Fatal("No estamos conectando bb")
		return
	}
	handlers.Handlers()
}
