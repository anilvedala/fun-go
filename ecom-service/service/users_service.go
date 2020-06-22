package service

import (
	"github.com/fun-go/ecom-service/database"
	"github.com/fun-go/ecom-service/dto/response"
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