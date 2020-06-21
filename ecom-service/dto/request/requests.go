package request

type UserCreationRequest struct {
	FirstName   string   `json:"first_name", required`
	SecondName  string   `json:"second_name", required`
	DateOfBirth string   `json:"date_of_birth",required`
	Email       string   `json:"email",required`
	PhoneNumber string   `json:"phone_number",required`
	Addresses   []string `json:"addresses"`
	BillingAddress string `json:"billing_address"`
}

type ItemRequest struct {
	Id string `json:"id", required`
	Name string `json:name,required`
	Quantity int `json:quantity,required,gt=0`
	SellerId   string `json:seller_id`
}

type OrderRequest struct {

	Items []ItemRequest `json:"items", required`
	DeliveryAddress string `json:"delivery_address", required`
	BillingAddress string `json:"billing_address", required`
}

