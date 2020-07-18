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

type UserCreationResponse struct {
	UserId string `json:"user_id"`
}

type OrderCreationResponse struct {
	OrderId   string `json:"order_id"`
	createdAt int64  `json:"created_at"`
	UserId    string `json:"user_id"`
}

type UserOrdersResponse struct {
	UserId string                  `json:"user_id"`
	Orders []OrderDetailsResponse `json:"orders"`
}

type OrderDetailsResponse struct {
	OrderId         string `json:"order_id"`
	Items           []ItemInOrderResponse `json:"items"`
	Status          string `json:"status"`
	DeliveryAddress string `json:"delivery_address"`
	BillingAddress  string `json:"billing_address"`
}

type ItemInOrderResponse struct {
	ItemId   string `json:"item_id"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
}
