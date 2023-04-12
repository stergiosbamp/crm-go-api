package controllers

import (
	"fmt"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/dao"
	"github.com/stergiosbamp/go-api/models"
)

var contactDAO = dao.NewContactDAO()

type ContactRequest struct {
	ID          uint   `json:"id" binding:"numeric"`
	ContactType string `json:"contactType" binding:"required"`
	FirstName   string `json:"firstName" binding:"required,max=50"`
	LastName    string `json:"lastName" binding:"required,max=50"`
	NickName    string `json:"nickName" binding:"required,max=50"`
	Gender      string `json:"gender" binding:"required"`
	Birthday    string `json:"birthday" binding:"required" time_format:"2006-01-02"`
	Language    string `json:"language" binding:"required,bcp47_language_tag"`
	JobTitle    string `json:"jobTitle" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Skype       string `json:"skype" binding:"required,max=25"`
	PhoneDirect string `json:"phoneDirect" binding:"required,numeric"`
	PhoneOffice string `json:"phoneOffice" binding:"required,numeric"`
	Mobile      string `json:"mobile" binding:"required,e164"`
	Notes       string `json:"notes" binding:"required"`
	CustomerID  uint   `json:"customerId" binding:"required,numeric"`
	AddressID   *uint  `json:"addressId" binding:"numeric"`
}

func GetContact(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uri.ID
	contact, err := contactDAO.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Contact with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, contact)
}

func GetContacts(ctx *gin.Context) {
	contacts, err := contactDAO.GetList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

func CreateContact(ctx *gin.Context) {
	var contactReq ContactRequest

	if err := ctx.ShouldBindJSON(&contactReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Shouldn't assign an address id to one that refers to a contact
	addressId := contactReq.AddressID
	address, err := addressDAO.GetById(*addressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", *addressId)})
		return
	}

	if address.Type != "contact" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v refers to a customer's address and not to a contact", addressId)})
		return
	}

	// Create a Model from request data
	contact := models.Contact{
		ContactType: contactReq.ContactType,
		FirstName:   contactReq.FirstName,
		LastName:    contactReq.LastName,
		NickName:    contactReq.NickName,
		Gender:      contactReq.Gender,
		Birthday:    contactReq.Birthday,
		Language:    contactReq.Language,
		JobTitle:    contactReq.JobTitle,
		Email:       contactReq.Email,
		Skype:       contactReq.Skype,
		PhoneDirect: contactReq.PhoneDirect,
		PhoneOffice: contactReq.PhoneOffice,
		Mobile:      contactReq.Mobile,
		Notes:       contactReq.Notes,
		CustomerID:  contactReq.CustomerID,
		AddressID:   contactReq.AddressID,
	}

	contactCreated, err := contactDAO.Create(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, contactCreated)
}

func UpdateContact(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contactReq ContactRequest

	if err := ctx.ShouldBindJSON(&contactReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if contact to update exists
	_, err := contactDAO.GetById(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Contact with id: %v doesn't exist", uri.ID)})
		return		
	}
	
	// Shouldn't assign an address id to one that refers to a contact
	addressId := contactReq.AddressID
	address, err := addressDAO.GetById(*addressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", *addressId)})
		return
	}

	if address.Type != "contact" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v refers to a customer's address and not to a contact", *addressId)})
		return
	}

	// Create a Model from request data
	contact := models.Contact{
		ID:          uri.ID,
		ContactType: contactReq.ContactType,
		FirstName:   contactReq.FirstName,
		LastName:    contactReq.LastName,
		NickName:    contactReq.NickName,
		Gender:      contactReq.Gender,
		Birthday:    contactReq.Birthday,
		Language:    contactReq.Language,
		JobTitle:    contactReq.JobTitle,
		Email:       contactReq.Email,
		Skype:       contactReq.Skype,
		PhoneDirect: contactReq.PhoneDirect,
		PhoneOffice: contactReq.PhoneOffice,
		Mobile:      contactReq.Mobile,
		Notes:       contactReq.Notes,
		CustomerID:  contactReq.CustomerID,
		AddressID:   contactReq.AddressID,
	}

	contactUpdated, err := contactDAO.Update(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, contactUpdated)
}

func DeleteContact(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errDeleted := contactDAO.Delete(uri.ID)
	if errDeleted != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", errDeleted.Error())})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"": ""})
}
