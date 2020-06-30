package routes

import (
	"github.com/fun-go/ecom-service/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes()  {


	// create a router with the basic router with all the basic middleware : logging, gzip
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	// GET all items
	router.GET("/items", controller.HandleGetAllItems)


	// User
	router.POST("/users", controller.HandleCreateUser)

	router.GET("/users/:userId", controller.HandleGetUser)

	//Order
	router.POST("/orders", controller.HandleCreateOrder)

	// running the server at the port number 8080
	router.Run(":8080")


}