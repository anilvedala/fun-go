package mapper

import (
	"errors"
	"github.com/fun-go/ecom-service/dto/request"
	"github.com/fun-go/ecom-service/models"
	"github.com/fun-go/ecom-service/utils"
	"github.com/google/uuid"
)

func MapUserRequestToModel(request *request.UserCreationRequest) (models.User, error) {

	if request == nil {
		return models.User{}, errors.New("received empty request")
	}

	dateOfBirth, err := utils.ParseCalendarDate(request.DateOfBirth)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:             uuid.New().String(),
		FirstName:      request.FirstName,
		SecondName:     request.SecondName,
		DateOfBirth:    dateOfBirth,
		Email:          request.Email,
		PhoneNumber:    request.PhoneNumber,
		Addresses:      request.Addresses,
		BillingAddress: request.BillingAddress,
		Orders:         make([]string, 16),
	}, nil

}


