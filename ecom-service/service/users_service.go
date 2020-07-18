package service

import (
	"github.com/fun-go/ecom-service/database"
	"github.com/fun-go/ecom-service/dto/request"
	"github.com/fun-go/ecom-service/dto/response"
	"github.com/fun-go/ecom-service/mapper"
)

func GetUserInformation(userId string) (*response.UserResponse, error) {

	user := database.GetUserDetails(userId)

	if user.Id == "" {
		return nil, nil
	}

	userResponse := &response.UserResponse{
		Id:             user.Id,
		FirstName:      user.FirstName,
		SecondName:     user.SecondName,
		DateOfBirth:    user.DateOfBirth,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		Addresses:      user.Addresses,
		BillingAddress: user.BillingAddress,
	}

	return userResponse, nil

}

func CreateUser(request *request.UserCreationRequest) (*response.UserCreationResponse, error) {

	userModel, err := mapper.MapUserRequestToModel(request)

	if err != nil {
		return nil, err
	}

	err = database.SaveUserDetails(userModel);
	if err != nil {
		return nil, err
	}

	return &response.UserCreationResponse{UserId: userModel.Id}, nil

}

func GetOrdersOfUser(userId string) (*response.UserOrdersResponse, error) {

	user := database.GetUserDetails(userId)
	if user.Id == "" {
		return nil, nil
	}

	orderIds := user.Orders

	orderDetails := database.GetOrderDetails(orderIds)

	return &response.UserOrdersResponse{
		UserId: userId,
		Orders: mapper.MapOrderModelToOrderResponse(orderDetails),
	}, nil


}
