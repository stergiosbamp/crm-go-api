package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/controllers"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	// separate group for auth to avoid auth middleware
	authGroup := rg.Group("/auth")
	
	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)
	// authGroup.POST("/logout", controllers.Logout)
}
