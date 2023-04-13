package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/dao"
	"github.com/stergiosbamp/go-api/models"
)

var addressDAO = dao.NewAddressDAO()

type AddressRequest struct {
	ID              uint    `json:"id" binding:"numeric"`
	CustomerID      *uint   `json:"customerId" binding:"excluded_unless=Type contact"` // Pointer since it can be null when referring to contacts' addresses, otherwise it defaults to customerId: 0
	Type            string  `json:"type" binding:"required,oneof=legal branch contact"`
	Address         string  `json:"address" binding:"required"`
	Pobox           string  `json:"pobox" binding:"required,numeric"`
	PostalCode      string  `json:"postalCode" binding:"required,numeric"`
	City            string  `json:"city" binding:"required"`
	Province        string  `json:"province" binding:"required"`
	Country         string  `json:"country" binding:"required"`
	AttentionPerson *string `json:"attentionPerson"`
}

type AddressResponse struct {
	ID              uint    `json:"id"`
	CustomerID      *uint   `json:"customerId,omitempty"`
	Type            string  `json:"type"`
	Address         string  `json:"address"`
	Pobox           string  `json:"pobox"`
	PostalCode      string  `json:"postalCode"`
	City            string  `json:"city"`
	Province        string  `json:"province"`
	Country         string  `json:"country"`
	AttentionPerson *string `json:"attentionPerson,omitempty"`
}

func GetAddress(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	id := uri.ID
	address, err := addressDAO.GetById(id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func GetAddresses(ctx *gin.Context) {
	addresses, err := addressDAO.GetList()
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func CreateAddress(ctx *gin.Context) {
	var addressReq AddressRequest
	var address models.Address

	if err := ctx.ShouldBindJSON(&addressReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if payload contains a customerId. If not it's meant for a contact.
	if addressReq.CustomerID != nil {
		
		// Customer exists?
		customerId := addressReq.CustomerID
		customer, err := customerDAO.GetById(*customerId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v doesn't exist", *customerId)})
			return
		}

		// We need validation only for "legal" address. Branch addresses must be created whether a customer is active or inactive.
		if addressReq.Type == "legal" {

			// Is customer active?
			if !(customer.Active) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v is not active", *customerId)})
				return
			}

			// Does customer already have a legal address?
			_, err := addressDAO.FindAddress(*customerId, addressReq.Type)

			// If err is not nil, it means it failed to find an address, thus customer doesn't have an address.
			if err == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v has already a legal address.", *customerId)})
				return
			}
		}
		
		// Populate an Address model from request data
		address = models.Address{
			CustomerID: addressReq.CustomerID,
			Type: addressReq.Type,
			Address: addressReq.Address,
			Pobox: addressReq.Pobox,
			PostalCode: addressReq.PostalCode,
			City: addressReq.City,
			Province: addressReq.Province,
			Country: addressReq.Country,
			AttentionPerson: addressReq.AttentionPerson,
		}
		
	} else {
		// Address refers to contact so omit customer data
		address = models.Address {
			Type: addressReq.Type,
			Address: addressReq.Address,
			Pobox: addressReq.Pobox,
			PostalCode: addressReq.PostalCode,
			City: addressReq.City,
			Province: addressReq.Province,
			Country: addressReq.Country,
		}
	}

	addressCreated, err := addressDAO.Create(&address)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	// Can use a universal response due to 'omitempty' tags
	addressRes := AddressResponse {
		ID: addressCreated.ID,
		CustomerID: addressReq.CustomerID,
		Type: addressCreated.Type,
		Address: addressCreated.Address,
		Pobox: addressCreated.Pobox,
		PostalCode: addressCreated.PostalCode,
		City: addressCreated.City,
		Province: addressCreated.Province,
		Country: addressCreated.Country,
		AttentionPerson: addressCreated.AttentionPerson,
	}
	
	ctx.JSON(http.StatusCreated, addressRes)
}

// Update operation is easy to break the integrity of type of addresses
// between customers. That's why it doesn't allow changing types.
// If you want to change type, create a new one and delete the old one.
func UpdateAddress(ctx *gin.Context) {
	var uri URI
	var oldAddress models.Address
	var addressReq AddressRequest

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&addressReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldAddress, err := addressDAO.GetById(uri.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", uri.ID)})
		return
	}

	// Business logic: Do not allow changing types, because
	// if changing from a branch to legal is forbidden
	// if changing from legal to branch then client must also call API to deactivate customer
	if oldAddress.Type != addressReq.Type {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Can't change type of address from %v to %v", oldAddress.Type, addressReq.Type)})
		return
	}

	// Create an Address model from the updated request data
	address := models.Address{
		ID: uri.ID,
		CustomerID: addressReq.CustomerID,
		Type: addressReq.Type,
		Address: addressReq.Address,
		Pobox: addressReq.Pobox,
		PostalCode: addressReq.PostalCode,
		City: addressReq.City,
		Province: addressReq.Province,
		Country: addressReq.Country,
		AttentionPerson: addressReq.AttentionPerson,
	}

	updatedAddress, err := addressDAO.Update(&address)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}
	
	addressRes := AddressResponse {
		ID: updatedAddress.ID,
		CustomerID: addressReq.CustomerID,
		Type: updatedAddress.Type,
		Address: updatedAddress.Address,
		Pobox: updatedAddress.Pobox,
		PostalCode: updatedAddress.PostalCode,
		City: updatedAddress.City,
		Province: updatedAddress.Province,
		Country: updatedAddress.Country,
		AttentionPerson: updatedAddress.AttentionPerson,
	}

	ctx.JSON(http.StatusOK, addressRes)
}

func DeleteAddress(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := addressDAO.GetById(uri.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", uri.ID)})
		return
	}	
	
	// To delete a "legal" branch you must first set customer to inactive.
	if address.Type == "legal" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Cannot delete a legal branch. Set customer with id: %v to inactive first.", *address.CustomerID)})
		return
	}
	
	errDeleted := addressDAO.Delete(uri.ID)
	if errDeleted != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", errDeleted.Error())})
		return
	}
	
	ctx.JSON(http.StatusNoContent, gin.H{"":""})
}
