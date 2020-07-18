package mapper

import (
	"github.com/fun-go/ecom-service/dto/response"
	"github.com/fun-go/ecom-service/models"
)

var MapOrderModelToOrderResponse = func(orders []models.Order) []response.OrderDetailsResponse {

	numOrders := len(orders)
	if numOrders == 0 {
		return nil
	}

	ordersDetailsResponse := make([]response.OrderDetailsResponse, len(orders))

	for index, order := range orders {
		ordersDetailsResponse[index] = response.OrderDetailsResponse{
			OrderId:         order.OrderId,
			Items:           mapItemInOrderModelToResponse(order.ItemOrders),
			Status:          order.Status,
			DeliveryAddress: order.DeliveryAddress,
			BillingAddress:  order.BillingAddress,
		}
	}

	return ordersDetailsResponse
}

var mapItemInOrderModelToResponse = func(itemsInOrder []models.ItemOrder) []response.ItemInOrderResponse {
	itemsInOrderResponse := make([]response.ItemInOrderResponse, len(itemsInOrder))

	for index, item := range itemsInOrder {
		itemsInOrderResponse[index] = response.ItemInOrderResponse{
			ItemId:   item.Id,
			Quantity: item.Quantity,
			Name:     item.Name,
		}
	}

	return itemsInOrderResponse
}
