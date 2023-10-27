package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/controllers"
)

func RegisterAddressesRoutes(rg *gin.RouterGroup) {
	rg.GET("/addresses", controllers.GetAddresses)
	rg.GET("/addresses/:id", controllers.GetAddress)
	rg.POST("/addresses", controllers.CreateAddress)
	rg.PUT("/addresses/:id", controllers.UpdateAddress)
	rg.DELETE("/addresses/:id", controllers.DeleteAddress)
	// TODO: EXPORT
}
