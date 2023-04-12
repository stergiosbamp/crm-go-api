package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/controllers"
)

func RegisterContactsRoutes(rg *gin.RouterGroup) {
	rg.GET("/contacts", controllers.GetContacts)
	rg.GET("/contacts/:id", controllers.GetContact)
	rg.POST("/contacts", controllers.CreateContact)
	rg.PUT("/contacts/:id", controllers.UpdateContact)
	rg.DELETE("/contacts/:id", controllers.DeleteContact)
	// TODO: IMPORT
	// TODO: EXPORT
}
