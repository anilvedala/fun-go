package response

import "time"

type ItemsResponse struct {
	Items []ItemResponse `json:"items"`
}

type ItemResponse struct {
	Id              string           `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	DefaultSellerId string           `json:"default_seller_id"`
	SellerInfo      []SellerResponse `json:"seller_info"`
}

type SellerResponse struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
	Price  float32 `json:"price"`
}

type UserResponse struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"first_name",validate:"required"`
	SecondName     string   `json:"second_name",validate:"required"`
	DateOfBirth    time.Time   `json:"date_of_birth",validate:"required"`
	Email          string   `json:"email",validate:"required"`
	PhoneNumber    string   `json:"phone_number",validate:"required"`
	Addresses      []string `json:"addresses"`
	BillingAddress string   `json:"billing_address"`
	signUpDate     time.Time `json:"sign_up_date"`
}
