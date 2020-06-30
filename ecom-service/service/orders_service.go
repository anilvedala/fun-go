package service

import (
	"github.com/fun-go/ecom-service/database"
	"github.com/fun-go/ecom-service/dto/request"
	"github.com/fun-go/ecom-service/dto/response"
	"github.com/fun-go/ecom-service/mapper"
	"github.com/google/uuid"
)

var CreateOrder = func(orderCreationRequest *request.OrderCreationRequest) (*response.OrderCreationResponse, error) {

	orderIdUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	orderId := orderIdUUID.String()
	order, err := mapper.MapOrderRequestToModel(orderCreationRequest, orderId)

	if err != nil {
		return nil, err
	}
	database.SaveOrder(order)

	return &response.OrderCreationResponse{
		OrderId: order.OrderId,
		UserId:  order.UserId,
	}, nil
}
