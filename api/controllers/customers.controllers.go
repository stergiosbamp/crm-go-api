package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/dao"
	"github.com/stergiosbamp/go-api/models"
)

var customerDAO = dao.NewCustomerDAO()

type CustomerRequest struct {
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

type CustomerPatchRequest struct {
	Active             *bool   	 `json:"active" binding:"required,boolean"`
}

func GetCustomer(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	id := uri.ID
	customer, err := customerDAO.GetById(id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with ID %v not found.", id)})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func GetCustomers(ctx *gin.Context) {
	customers, err := customerDAO.GetList()
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func CreateCustomer(ctx *gin.Context) {
	var customerReq CustomerRequest

	if err := ctx.ShouldBindJSON(&customerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate a Model from request data
	customer := &models.Customer{
		Active: *customerReq.Active,
		Name: customerReq.Name,
		ShortName: customerReq.ShortName,
		VatNumber: customerReq.VatNumber,
		VatApply: *customerReq.VatApply,
		RegistrationNumber: customerReq.RegistrationNumber,
		DunsNumber: customerReq.DunsNumber,
		TaxExempt: *customerReq.TaxExempt,
		Language: customerReq.Language,
		Email: customerReq.Email,
		Phone: customerReq.Phone,
		Fax: customerReq.Fax,
	}
	
	customerCreated, err := customerDAO.Create(customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, customerCreated)
}

func UpdateCustomer(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var customerReq CustomerRequest

	if err := ctx.ShouldBindJSON(&customerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Check if customer to update exists
	_, err := customerDAO.GetById(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v doesn't exist", uri.ID)})
		return		
	}

	// Populate a Model with the id path parameter and the updated request data.
	// If in the future, PUT can work without requiring all fields, then use the "oldCustomer" data from the above operation of DAO 
	// in conjunction with the "customerReq" data to fill up the model for the missing fields.
	customer := &models.Customer{
		ID: uri.ID, // inject ID for update
		Active: *customerReq.Active,
		Name: customerReq.Name,
		ShortName: customerReq.ShortName,
		VatNumber: customerReq.VatNumber,
		VatApply: *customerReq.VatApply,
		RegistrationNumber: customerReq.RegistrationNumber,
		DunsNumber: customerReq.DunsNumber,
		TaxExempt: *customerReq.TaxExempt,
		Language: customerReq.Language,
		Email: customerReq.Email,
		Phone: customerReq.Phone,
		Fax: customerReq.Fax,
	}

	customerUpdated, err := customerDAO.Update(customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, customerUpdated)
}

func PatchCustomer(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var patchBody CustomerPatchRequest

	if err := ctx.ShouldBindJSON(&patchBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if customer to update exists
	oldCustomer, err := customerDAO.GetById(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v doesn't exist", uri.ID)})
		return		
	}

	// ADD BUSINESS LOGIC THAT CANNOT ACTIVATE A CUSTOMER WITHOUT HAVING IN DB A LEGAL ADDRESS.
	
	customerPatched := models.Customer{
		ID: uri.ID, // inject ID
		Active: *patchBody.Active, // update status
		Name: oldCustomer.Name,
		ShortName: oldCustomer.ShortName,
		VatNumber: oldCustomer.VatNumber,
		VatApply: oldCustomer.VatApply,
		RegistrationNumber: oldCustomer.RegistrationNumber,
		DunsNumber: oldCustomer.DunsNumber,
		TaxExempt: oldCustomer.TaxExempt,
		Language: oldCustomer.Language,
		Email: oldCustomer.Email,
		Phone: oldCustomer.Phone,
		Fax: oldCustomer.Fax,
	}
	
	newCustomer, err := customerDAO.Update(&customerPatched)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, newCustomer)
}

func DeleteCustomer(ctx *gin.Context) {
	var uri URI

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := customerDAO.GetById(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Customer with id: %v doesn't exist", uri.ID)})
		return		
	}

	errDeleted := customerDAO.Delete(uri.ID)
	if errDeleted != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"":""})
}
