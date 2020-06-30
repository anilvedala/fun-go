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
