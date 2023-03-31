package main

import (
	// "log"

	"github.com/stergiosbamp/go-api/database"
	"github.com/stergiosbamp/go-api/models"
)

func main() {
	db, _ := database.GetDB()
	db.AutoMigrate(&models.Customer{}, &models.Contact{}, &models.Address{})

}
