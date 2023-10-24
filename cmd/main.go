package main

import (
	"log"

	"github.com/stergiosbamp/go-api/src/database"
	"github.com/stergiosbamp/go-api/src/models"
	"github.com/stergiosbamp/go-api/src/routes"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Failed to start database", err)
	}

	err = db.AutoMigrate(&models.Customer{}, &models.Contact{}, &models.Address{}, &models.User{}, &models.Token{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	routes.InitRoutes()
}
