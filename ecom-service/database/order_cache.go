package database

import (
	"errors"
	"github.com/fun-go/ecom-service/models"
)

var (
	orders map[string]models.Order
)

func Init() {
	orders = make(map[string]models.Order)
}

var SaveOrder = func(order models.Order) error {

	if orders == nil {
		return errors.New("database is not initialised")
	}
	orders[order.OrderId] = order

	return nil
}

var GetOrderDetails = func(orderIds []string) []models.Order {

	orderDetails := make([]models.Order, len(orderIds))

	numOrdersFound := 0

	for _, orderId := range orderIds {
		if order, ok := orders[orderId]; ok {
			orderDetails[numOrdersFound] = order
			numOrdersFound++
		}
	}

	return orderDetails[:numOrdersFound]
}
