package controller

import (
	"github.com/fun-go/ecom-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var HandleGetAllItems = func(context *gin.Context){

	ItemsResponse, err := service.GetAllItemsInStore()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Error occurred in the server while processing the request.")
	}

	context.JSON(http.StatusOK, ItemsResponse)

}