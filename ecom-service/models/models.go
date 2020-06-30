package models

import (
	"time"
)

type Item struct {
	Id              string
	Name            string
	Description     string
	DefaultSellerId string
	SellerInfo      []ItemSeller
}

type ItemSeller struct {
	SellerId string
	Price    float32
}

type Seller struct {
	Id     string
	Name   string
	Rating float32
	Price  float32
}

type User struct {
	Id             string
	FirstName      string
	SecondName     string
	DateOfBirth    time.Time
	Email          string
	PhoneNumber    string
	Addresses      []string
	BillingAddress string
	Orders         []string
	signUpDate     time.Time
}

type Order struct {
	OrderId         string
	UserId          string
	DeliveryAddress string
	BillingAddress  string
	ItemOrders      []ItemOrder
}

type ItemOrder struct {
	Id string `json:"id",validate:"required"`
	Name string `json:"name",validate:"required"`
	Quantity int `json:"quantity",validate:"required"`
	SellerId   string `json:"seller_id"`
}
