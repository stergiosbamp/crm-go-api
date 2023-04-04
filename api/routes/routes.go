package routes

import "github.com/gin-gonic/gin"

func InitRoutes() {
	route := gin.Default()

	v1 := route.Group("v1")

	RegisterCustomersRoutes(v1)
	RegisterAddressesRoutes(v1)
	// REGISTER CONTACTS

	route.Run()
}
