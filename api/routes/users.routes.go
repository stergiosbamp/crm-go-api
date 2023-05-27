package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/controllers"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", controllers.Register)
	rg.POST("/login", controllers.Login)
	// rg.POST("/logout", controllers.CreateAddress)
}