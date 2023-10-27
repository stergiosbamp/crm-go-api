package controllers

import (
	"fmt"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/dao"
	"github.com/stergiosbamp/go-api/src/models"
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
	AddressID   *uint  `json:"addressId" binding:"omitempty,numeric"`
}

type ContactResponse struct {
	ID          uint   `json:"id"`
	ContactType string `json:"contactType"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NickName    string `json:"nickName"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	Language    string `json:"language"`
	JobTitle    string `json:"jobTitle"`
	Email       string `json:"email"`
	Skype       string `json:"skype"`
	PhoneDirect string `json:"phoneDirect"`
	PhoneOffice string `json:"phoneOffice"`
	Mobile      string `json:"mobile"`
	Notes       string `json:"notes"`
	CustomerID  uint   `json:"customerId"`
	AddressID   *uint  `json:"addressId,omitempty"`
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

	// Create a ContactResponse from the contact model
	contactRes := ContactResponse{
		ID:          contact.ID,
		ContactType: contact.ContactType,
		FirstName:   contact.FirstName,
		LastName:    contact.LastName,
		NickName:    contact.NickName,
		Gender:      contact.Gender,
		Birthday:    contact.Birthday,
		Language:    contact.Language,
		JobTitle:    contact.JobTitle,
		Email:       contact.Email,
		Skype:       contact.Skype,
		PhoneDirect: contact.PhoneDirect,
		PhoneOffice: contact.PhoneOffice,
		Mobile:      contact.Mobile,
		Notes:       contact.Notes,
		CustomerID:  contact.CustomerID,
		AddressID:   contact.AddressID,
	}

	ctx.JSON(http.StatusOK, contactRes)
}

func GetContacts(ctx *gin.Context) {
	contacts, err := contactDAO.GetList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	var contactsRes []ContactResponse

	for _, contact := range contacts {
		contact := ContactResponse{
			ID:          contact.ID,
			ContactType: contact.ContactType,
			FirstName:   contact.FirstName,
			LastName:    contact.LastName,
			NickName:    contact.NickName,
			Gender:      contact.Gender,
			Birthday:    contact.Birthday,
			Language:    contact.Language,
			JobTitle:    contact.JobTitle,
			Email:       contact.Email,
			Skype:       contact.Skype,
			PhoneDirect: contact.PhoneDirect,
			PhoneOffice: contact.PhoneOffice,
			Mobile:      contact.Mobile,
			Notes:       contact.Notes,
			CustomerID:  contact.CustomerID,
			AddressID:   contact.AddressID,
		}
		contactsRes = append(contactsRes, contact)
	}

	ctx.JSON(http.StatusOK, contactsRes)
}

func CreateContact(ctx *gin.Context) {
	var contactReq ContactRequest

	if err := ctx.ShouldBindJSON(&contactReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addressId := contactReq.AddressID
	// Address is optional for contacts
	if addressId != nil {
		// Shouldn't assign an address id to one that refers to a contact
		address, err := addressDAO.GetById(*addressId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v doesn't exist", *addressId)})
			return
		}
		if address.Type != "contact" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address with id: %v refers to a customer's address and not to a contact", *addressId)})
			return
		}
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

	contactRes := ContactResponse{
		ID:          contactCreated.ID,
		ContactType: contactCreated.ContactType,
		FirstName:   contactCreated.FirstName,
		LastName:    contactCreated.LastName,
		NickName:    contactCreated.NickName,
		Gender:      contactCreated.Gender,
		Birthday:    contactCreated.Birthday,
		Language:    contactCreated.Language,
		JobTitle:    contactCreated.JobTitle,
		Email:       contactCreated.Email,
		Skype:       contactCreated.Skype,
		PhoneDirect: contactCreated.PhoneDirect,
		PhoneOffice: contactCreated.PhoneOffice,
		Mobile:      contactCreated.Mobile,
		Notes:       contactCreated.Notes,
		CustomerID:  contactCreated.CustomerID,
		AddressID:   contactCreated.AddressID,
	}

	ctx.JSON(http.StatusCreated, contactRes)
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

	contactRes := ContactResponse{
		ID:          contactUpdated.ID,
		ContactType: contactUpdated.ContactType,
		FirstName:   contactUpdated.FirstName,
		LastName:    contactUpdated.LastName,
		NickName:    contactUpdated.NickName,
		Gender:      contactUpdated.Gender,
		Birthday:    contactUpdated.Birthday,
		Language:    contactUpdated.Language,
		JobTitle:    contactUpdated.JobTitle,
		Email:       contactUpdated.Email,
		Skype:       contactUpdated.Skype,
		PhoneDirect: contactUpdated.PhoneDirect,
		PhoneOffice: contactUpdated.PhoneOffice,
		Mobile:      contactUpdated.Mobile,
		Notes:       contactUpdated.Notes,
		CustomerID:  contactUpdated.CustomerID,
		AddressID:   contactUpdated.AddressID,
	}

	ctx.JSON(http.StatusOK, contactRes)
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
