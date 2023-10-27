package main

import (
	"log"

	"github.com/stergiosbamp/go-api/src/database"
	"github.com/stergiosbamp/go-api/src/routes"
)

func main() {
	// Migrate database
	err := database.Migrate()
	if err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	// Initialize API routes
	routes.Init()
}
