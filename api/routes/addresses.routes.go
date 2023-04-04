package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/controllers"
)


func RegisterAddressesRoutes(rg *gin.RouterGroup) {
	rg.GET("/addresses", controllers.GetAddresses)
	rg.GET("/addresses/:id", controllers.GetAddress)
	rg.POST("/addresses", controllers.CreateAddress)
	// rg.PUT("/addresses/:id", controllers.UpdateOrCreateCustomer)
	// rg.PATCH("/addresses/:id", controllers.PatchCustomer)
	// rg.DELETE("/addresses/:id", controllers.DeleteCustomer)
	// EXPORT
}