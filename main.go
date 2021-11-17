package main

import (
	"log"

	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/handler"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("No connection")
		return
	}
	handler.RoutesHandler()
}
