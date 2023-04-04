package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	ID                 	uint  	 `json:"id" binding:"numeric"`
	CustomerID         	*uint    `json:"customerId" binding:"excluded_unless=Type contact"`  // Pointer since it can be null when refereing to contacts' addresses, otherwise it defaults to customerId: 0
	Type				string	 `json:"type" binding:"required,oneof=legal branch contact"`
	Address         	string	 `json:"address" binding:"required"`
	Pobox           	string	 `json:"pobox" binding:"required,numeric"`
	PostalCode      	string	 `json:"postalCode" binding:"required,numeric"`
	City            	string	 `json:"city" binding:"required"`
	Province        	string   `json:"province" binding:"required"`
	Country         	string	 `json:"country" binding:"required"`
	AttentionPerson 	*string	 `json:"attentionPerson"`
}

func GetAddress(ctx *gin.Context) {
	var uri URI
	var address Address

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uri.ID
	
	res := db.Where("id = ?", id).First(&address)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func GetAddresses(ctx *gin.Context) {
	var addresses []Address
	
	res := db.Find(&addresses)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func CreateAddress(ctx *gin.Context) {
	var address Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	res := db.Create(&address)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusCreated, address)
}

func UpdateOrCreateAddress(ctx *gin.Context) {
	var uri URI
	var updatedAdress Address

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := ctx.ShouldBindJSON(&updatedAdress); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the ID for Update or Create
	updatedAdress.ID = uri.ID

	res := db.Save(&updatedAdress)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, updatedAdress)
}

func DeleteAddress(ctx *gin.Context) {
	var uri URI
	var address Address

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uri.ID
	
	res := db.Where("id = ?", id).Unscoped().Delete(&address)

	// ADD LOGIC WHEN DELETING A 'LEGAL' BRANCH TO CHANGE (OR NOT) THE CUSTOMER 'active' STATUS.
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusNoContent, address)
}
