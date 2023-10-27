package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/config"
	"github.com/stergiosbamp/go-api/src/middleware"
)

var conf config.Config

func Init() {
	route := gin.Default()

	// API info, aliveness.
	route.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "A pluggable, simple and fast CRM API")
	})

	v1 := route.Group("v1")

	RegisterUserRoutes(v1)

	if conf.UseAuth() {
		// all endpoints below require authentication
		v1.Use(middleware.JwtAuth())
	}

	RegisterCustomersRoutes(v1)
	RegisterAddressesRoutes(v1)
	RegisterContactsRoutes(v1)

	route.Run()
}
