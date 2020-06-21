package response


type ItemsResponse struct {
	Items []ItemResponse `json:"items"`
}

type ItemResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	DefaultSellerId string `json:"default_seller_id"`
	SellerInfo []SellerResponse `json:"seller_info"`
}

type SellerResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Rating float32 `json:"rating"`
	Price float32 `json:"price"`
}
