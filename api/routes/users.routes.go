package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/controllers"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	// group for only auth routes
	authGroup := rg.Group("/auth")
	
	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)
	// rg.POST("/logout", controllers.CreateAddress)
}