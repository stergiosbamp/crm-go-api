package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/dao"
	"github.com/stergiosbamp/go-api/src/models"
)

var customerDAO = dao.NewCustomerDAO()

type CustomerRequest struct {
	ID                 uint   `json:"id" binding:"numeric"`
	Active             *bool  `json:"active" binding:"required,boolean"`
	Name               string `json:"name" binding:"required,max=100"`
	ShortName          string `json:"shortName" binding:"required,max=50"`
	VatNumber          string `json:"vatNumber" binding:"required,numeric"`
	VatApply           *bool  `json:"vatApply" binding:"boolean"`
	RegistrationNumber string `json:"registrationNumber" binding:"required,numeric"`
	DunsNumber         string `json:"dunsNumber" binding:"required,numeric"`
	TaxExempt          *bool  `json:"taxExempt" binding:"boolean"`
	Language           string `json:"language" binding:"required,bcp47_language_tag"`
	Email              string `json:"email" binding:"required,email"`
	Phone              string `json:"phone" binding:"required,numeric"`
	Fax                string `json:"fax" binding:"required"`
}

// Different struct for response instead of embedding json tags into the model and return model.
// 1. separation of concerns
// 2. easily populate model fields (e.g. Customer) without requiring to be pointer which is hard to retrieve embedded data.
// 3. future-proofing to involve API independently without affecting underlying model
type CustomerResponse struct {
	ID                 uint   `json:"id"`
	Active             bool   `json:"active"`
	Name               string `json:"name"`
	ShortName          string `json:"shortName"`
	VatNumber          string `json:"vatNumber"`
	VatApply           bool   `json:"vatApply"`
	RegistrationNumber string `json:"registrationNumber"`
	DunsNumber         string `json:"dunsNumber"`
	TaxExempt          bool   `json:"taxExempt"`
	Language           string `json:"language"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Fax                string `json:"fax"`
}

type URI struct {
	ID uint `uri:"id" binding:"required"`
}

type CustomerPatchRequest struct {
	Active *bool `json:"active" binding:"required,boolean"`
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
	var customerRes CustomerResponse

	if err := ctx.ShouldBindJSON(&customerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate a Model from request data
	customer := &models.Customer{
		Active:             *customerReq.Active,
		Name:               customerReq.Name,
		ShortName:          customerReq.ShortName,
		VatNumber:          customerReq.VatNumber,
		VatApply:           *customerReq.VatApply,
		RegistrationNumber: customerReq.RegistrationNumber,
		DunsNumber:         customerReq.DunsNumber,
		TaxExempt:          *customerReq.TaxExempt,
		Language:           customerReq.Language,
		Email:              customerReq.Email,
		Phone:              customerReq.Phone,
		Fax:                customerReq.Fax,
	}

	customerCreated, err := customerDAO.Create(customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	// Map model to response struct
	customerRes = CustomerResponse{
		ID:                 customerCreated.ID,
		Active:             customerCreated.Active,
		Name:               customerCreated.Name,
		ShortName:          customerCreated.ShortName,
		VatNumber:          customerCreated.VatNumber,
		VatApply:           customerCreated.VatApply,
		RegistrationNumber: customerCreated.RegistrationNumber,
		DunsNumber:         customerCreated.DunsNumber,
		TaxExempt:          customerCreated.TaxExempt,
		Language:           customerCreated.Language,
		Email:              customerCreated.Email,
		Phone:              customerCreated.Phone,
		Fax:                customerCreated.Fax,
	}

	ctx.JSON(http.StatusCreated, customerRes)
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
		ID:                 uri.ID, // inject ID for update
		Active:             *customerReq.Active,
		Name:               customerReq.Name,
		ShortName:          customerReq.ShortName,
		VatNumber:          customerReq.VatNumber,
		VatApply:           *customerReq.VatApply,
		RegistrationNumber: customerReq.RegistrationNumber,
		DunsNumber:         customerReq.DunsNumber,
		TaxExempt:          *customerReq.TaxExempt,
		Language:           customerReq.Language,
		Email:              customerReq.Email,
		Phone:              customerReq.Phone,
		Fax:                customerReq.Fax,
	}

	customerUpdated, err := customerDAO.Update(customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	customerRes := CustomerResponse{
		ID:                 customerUpdated.ID,
		Active:             customerUpdated.Active,
		Name:               customerUpdated.Name,
		ShortName:          customerUpdated.ShortName,
		VatNumber:          customerUpdated.VatNumber,
		VatApply:           customerUpdated.VatApply,
		RegistrationNumber: customerUpdated.RegistrationNumber,
		DunsNumber:         customerUpdated.DunsNumber,
		TaxExempt:          customerUpdated.TaxExempt,
		Language:           customerUpdated.Language,
		Email:              customerUpdated.Email,
		Phone:              customerUpdated.Phone,
		Fax:                customerUpdated.Fax,
	}

	ctx.JSON(http.StatusOK, customerRes)
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

	customer := models.Customer{
		ID:                 uri.ID,            // inject ID
		Active:             *patchBody.Active, // update status
		Name:               oldCustomer.Name,
		ShortName:          oldCustomer.ShortName,
		VatNumber:          oldCustomer.VatNumber,
		VatApply:           oldCustomer.VatApply,
		RegistrationNumber: oldCustomer.RegistrationNumber,
		DunsNumber:         oldCustomer.DunsNumber,
		TaxExempt:          oldCustomer.TaxExempt,
		Language:           oldCustomer.Language,
		Email:              oldCustomer.Email,
		Phone:              oldCustomer.Phone,
		Fax:                oldCustomer.Fax,
	}

	customerPatched, err := customerDAO.Update(&customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error DB. Full error %v", err.Error())})
		return
	}

	customerRes := CustomerResponse{
		ID:                 customerPatched.ID,
		Active:             customerPatched.Active,
		Name:               customerPatched.Name,
		ShortName:          customerPatched.ShortName,
		VatNumber:          customerPatched.VatNumber,
		VatApply:           customerPatched.VatApply,
		RegistrationNumber: customerPatched.RegistrationNumber,
		DunsNumber:         customerPatched.DunsNumber,
		TaxExempt:          customerPatched.TaxExempt,
		Language:           customerPatched.Language,
		Email:              customerPatched.Email,
		Phone:              customerPatched.Phone,
		Fax:                customerPatched.Fax,
	}

	ctx.JSON(http.StatusOK, customerRes)
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
	ctx.JSON(http.StatusNoContent, gin.H{"": ""})
}
