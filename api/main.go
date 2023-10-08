package main

import (
	"github.com/stergiosbamp/go-api/database"
	"github.com/stergiosbamp/go-api/models"
	"github.com/stergiosbamp/go-api/routes"
)

func main() {
	db, _ := database.GetDB()
	db.AutoMigrate(&models.Customer{}, &models.Contact{}, &models.Address{}, &models.User{}, &models.Token{})
	
	routes.InitRoutes()

}
