package models


type Item struct {
	Id string
	Name string
	Description string
	DefaultSellerId string
	SellerInfo []ItemSeller
}


type ItemSeller struct {
	SellerId string
	Price float32
}

type Seller struct {
	Id string
	Name string
	Rating float32
	Price float32
}