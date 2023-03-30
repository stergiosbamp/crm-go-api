package main

import (
	"log"

	"github.com/stergiosbamp/go-api/models"
	"github.com/stergiosbamp/go-api/database"
	
)


func main() {
	db, _ := database.GetDB()

	var customer models.Customer
	db.First(&customer)
	
	log.Println(customer)
}
