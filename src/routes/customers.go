package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/stergiosbamp/go-api/src/controllers"
)

// Routes function to serve endpoints
func RegisterCustomersRoutes(rg *gin.RouterGroup) {
	rg.GET("/customers", controllers.GetCustomers)
	rg.GET("/customers/:id", controllers.GetCustomer)
	rg.POST("/customers", controllers.CreateCustomer)
	rg.PUT("/customers/:id", controllers.UpdateCustomer)
	rg.PATCH("/customers/:id", controllers.PatchCustomer)
	rg.DELETE("/customers/:id", controllers.DeleteCustomer)
	// TODO: IMPORT
	// TODO: EXPORT
}
