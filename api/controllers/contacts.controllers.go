package controllers

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID                 	uint  	 `json:"id" binding:"numeric"`
	ContactType 		string 	 `json:"contactType" binding:"required"`
	FirstName   		string   `json:"firstName" binding:"required,max=50"`
	LastName    		string   `json:"lastName" binding:"required,max=50"`
	NickName    		string   `json:"nickName" binding:"required,max=50"`
	Gender      		string   `json:"gender" binding:"required"`
	Birthday    		string	 `json:"birthday" binding:"required" time_format:"2006-01-02"`
	Language    		string   `json:"language" binding:"required,bcp47_language_tag"`
	JobTitle    		string   `json:"jobTitle" binding:"required"`
	Email       		string   `json:"email" binding:"required,email"`
	Skype       		string   `json:"skype" binding:"required,max=25"`
	PhoneDirect 		string   `json:"phoneDirect" binding:"required,numeric"`
	PhoneOffice 		string   `json:"phoneOffice" binding:"required,numeric"`
	Mobile      		string   `json:"mobile" binding:"required,e164"`
	Notes       		string   `json:"notes" binding:"required"`
	CustomerID  		uint     `json:"customerId" binding:"required,numeric"`
	//Customer 			Customer 
	AddressID   		uint     `json:"addressId" binding:"required,numeric"`
	//Address 			Address 
}

func GetContact(ctx *gin.Context) {
	var uri URI
	var contact Contact

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uri.ID
	
	res := db.Where("id = ?", id).First(&contact)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Contact with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, contact)
}

func GetContacts(ctx *gin.Context) {
	var contacts []Contact
	
	res := db.Find(&contacts)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

func CreateContact(ctx *gin.Context) {
	var contact Contact

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	res := db.Create(&contact)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusCreated, contact)
}

func UpdateOrCreateContact(ctx *gin.Context) {
	var uri URI
	var updatedContact Contact

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := ctx.ShouldBindJSON(&updatedContact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the ID for Update or Create
	updatedContact.ID = uri.ID

	res := db.Save(&updatedContact)
	
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, updatedContact)
}

func DeleteContact(ctx *gin.Context) {
	var uri URI
	var contact Contact

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uri.ID
	
	res := db.Where("id = ?", id).Unscoped().Delete(&contact)

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", res.Error)})
		return
	}

	ctx.JSON(http.StatusNoContent, contact)
}
