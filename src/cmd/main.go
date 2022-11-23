package main

import (
	"log"
	"net/http"

	"github.com/franlauriano/docker-compose/cmd/controller"
	"github.com/franlauriano/docker-compose/lib/database"
)

func main() {
	// Connect to database
	database.OpenPostgres("beach", "password", "beach", "db", 5432)
	defer database.Close()

	// Routes
	http.HandleFunc("/beaches", controller.BeachHandle)

	// Start http server
	addr := ":8080"
	log.Printf("Listening on http://%s", addr)
	http.ListenAndServe(addr, nil)
}
