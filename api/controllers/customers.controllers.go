package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/database"
)

var db, _ = database.GetDB()

type Customer struct {
	ID                 uint  	 `json:"id" binding:"numeric"`
	Active             *bool   	 `json:"active" binding:"required,boolean"`
	Name               string 	 `json:"name" binding:"required,max=100"`
	ShortName          string	 `json:"shortName" binding:"required,max=50"`
	VatNumber          string	 `json:"vatNumber" binding:"required,numeric"`
	VatApply           *bool	 `json:"vatApply" binding:"boolean"`
	RegistrationNumber string	 `json:"registrationNumber" binding:"required,numeric"`
	DunsNumber         string	 `json:"dunsNumber" binding:"required,numeric"`
	TaxExempt          *bool   	 `json:"taxExempt" binding:"boolean"`
	Language           string	 `json:"language" binding:"required,bcp47_language_tag"`
	Email              string	 `json:"email" binding:"required,email"`
	Phone              string	 `json:"phone" binding:"required,numeric"`
	Fax                string	 `json:"fax" binding:"required"`
}

type URI struct {
	ID  uint `uri:"id" binding:"required"`
}

func GetCustomer(ctx *gin.Context) {
	var uri URI
	var customer Customer

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uri.ID
	
	res := db.Where("id = ?", id).First(&customer)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func GetCustomers(ctx *gin.Context) {
	var customers []Customer
	
	res := db.Find(&customers)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func CreateCustomer(ctx *gin.Context) {
	var customer Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	res := db.Create(&customer)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(ctx *gin.Context) {
	var uri URI
	var updatedCustomer Customer

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := ctx.ShouldBindJSON(&updatedCustomer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the ID for Update or Create
	updatedCustomer.ID = uri.ID

	res := db.Save(&updatedCustomer)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, updatedCustomer)
}

func DeleteCustomer(ctx *gin.Context) {
	var uri URI
	var customer Customer

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uri.ID

	res := db.Where("id = ?", id).Unscoped().Delete(&customer)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusNoContent, customer)
}
