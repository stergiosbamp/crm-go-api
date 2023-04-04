package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/stergiosbamp/go-api/controllers"
)

// Routes function to serve endpoints
func RegisterRoutes() {
	route := gin.Default()

	v1 := route.Group("v1")

	v1.GET("/customers", controllers.GetCustomers)
	v1.GET("/customers/:id", controllers.GetCustomer)
	v1.POST("/customers", controllers.CreateCustomer)
	v1.PUT("/customers/:id", controllers.UpdateOrCreateCustomer)
	v1.PATCH("/customers/:id", controllers.PatchCustomer)
	v1.DELETE("/customers/:id", controllers.DeleteCustomer)
	// TODO: IMPORT
	// TODO: EXPORT

	route.Run()
}
