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
	signUpDate  time.Time
}
