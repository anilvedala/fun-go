package controller

import (
	"github.com/fun-go/ecom-service/constants"
	errors "github.com/fun-go/ecom-service/dto/error"
	"github.com/fun-go/ecom-service/dto/request"
	"github.com/fun-go/ecom-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var HandleGetAllItems = func(context *gin.Context) {

	itemsResponse, err := service.GetAllItemsInStore()
	if err != nil {
		context.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			ErrorMessage: constants.InternalErrorMessage,
		})
	}

	context.JSON(http.StatusOK, itemsResponse)

}

var HandleGetUser = func(context *gin.Context) {
	userId := context.Param(constants.UserIdPathParamKey)
	usersResponse, err := service.GetUserInformation(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			ErrorMessage: constants.InternalErrorMessage,
		})
	}

	if usersResponse == nil {
		context.JSON(http.StatusNotFound, nil)
	}

	context.JSON(http.StatusOK, usersResponse)
}

var HandleCreateUser = func(context *gin.Context) {

	var userCreationRequest *request.UserCreationRequest
	if err := context.ShouldBindJSON(&userCreationRequest); err != nil {
		context.JSON(http.StatusBadRequest, errors.ErrorResponse{
			ErrorMessage: constants.DeserializationErrorMessage,
		})
	}

	userCreationResponse, err := service.CreateUser(userCreationRequest)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errors.ErrorResponse{ErrorMessage: err.Error()})
	}

	context.JSON(http.StatusCreated, userCreationResponse)

}
