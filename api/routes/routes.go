package routes

import "github.com/gin-gonic/gin"

func InitRoutes() {
	route := gin.Default()

	// API info, aliveness.
	route.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "A pluggable, simple and fast CRM API")
	})

	v1 := route.Group("v1")

	RegisterCustomersRoutes(v1)
	RegisterAddressesRoutes(v1)
	RegisterContactsRoutes(v1)
	RegisterUserRoutes(v1)

	route.Run()
}
