package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	ID              uint    `json:"id" binding:"numeric"`
	CustomerID      *uint   `json:"customerId" binding:"excluded_unless=Type contact"` // Pointer since it can be null when refereing to contacts' addresses, otherwise it defaults to customerId: 0
	Type            string  `json:"type" binding:"required,oneof=legal branch contact"`
	Address         string  `json:"address" binding:"required"`
	Pobox           string  `json:"pobox" binding:"required,numeric"`
	PostalCode      string  `json:"postalCode" binding:"required,numeric"`
	City            string  `json:"city" binding:"required"`
	Province        string  `json:"province" binding:"required"`
	Country         string  `json:"country" binding:"required"`
	AttentionPerson *string `json:"attentionPerson"`
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

	// Check if payload contains a customerId. If not it's meant for a contact.
	if address.CustomerID != nil {

		customerId := address.CustomerID
		// Customer exists?
		var customer Customer
		res := db.First(&customer, customerId)
		if res.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v doesn't exist", *customerId)})
			return
		}

		// We need validation only for "legal" address. Branch addresses must be created whether a customer is active or inactive.
		if address.Type == "legal" {

			// Is customer active?
			if !*(customer.Active) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v is not active", *customerId)})
				return
			}

			// Does customer already have a legal address?
			var legalAddress Address
			res := db.Where("customer_id = ? AND type = ?", customerId, "legal").First(&legalAddress)
			if res.Error == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v has already a legal address.", *customerId)})
				return
			}
		}
		// set the customer ID for the address
		address.CustomerID = customerId
	}

	resCreate := db.Create(&address)

	if resCreate.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", resCreate.Error)})
		return
	}

	ctx.JSON(http.StatusCreated, address)
}

func UpdateAddress(ctx *gin.Context) {
	// Update operation is easy to break the integrity of type of addresses
	// between customers. That's why it doesn't allow changing types.
	// If you want to change type, create a new one and delete the old one.

	var uri URI
	var oldAddress Address
	var updatedAdress Address

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&updatedAdress); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := db.First(&oldAddress, uri.ID)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", uri.ID)})
		return
	}

	if oldAddress.Type != updatedAdress.Type {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Can't change type of address from %v to %v", oldAddress.Type, updatedAdress.Type)})
		return
	}

	updatedAdress.ID = uri.ID
	resUpdate := db.Save(&updatedAdress)

	if resUpdate.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", resUpdate.Error.Error())})
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
