package request

type UserCreationRequest struct {
	FirstName   string   `json:"first_name",validate:"required"`
	SecondName  string   `json:"second_name",validate:"required"`
	DateOfBirth string   `json:"date_of_birth",validate:"required"`
	Email       string   `json:"email",validate:"required"`
	PhoneNumber string   `json:"phone_number",validate:"required"`
	Addresses   []string `json:"addresses"`
	BillingAddress string `json:"billing_address"`
}

type ItemRequest struct {
	Id string `json:"id",validate:"required"`
	Name string `json:"name",validate:"required"`
	Quantity int `json:"quantity",validate:"required"`
	SellerId   string `json:"seller_id"`
}

type OrderRequest struct {

	Items []ItemRequest `json:"items"`
	DeliveryAddress string `json:"delivery_address"`
	BillingAddress string `json:"billing_address"`
}

