package service

import (
	"github.com/fun-go/ecom-service/database"
	"github.com/fun-go/ecom-service/dto/response"
)

func GetAllItemsInStore() (*response.ItemsResponse, error) {

	items := database.GetAllItemsInStore()
	sellers := database.GetAllSellers()

	itemsArray := make([]response.ItemResponse, len(items))

	for _, item := range items {
		itemSellersResponse := make([]response.SellerResponse, len(item.SellerInfo))

		for _, sellerInfo := range item.SellerInfo {
			sellerResponse := response.SellerResponse{
				Id:     sellerInfo.SellerId,
				Name:   sellers[sellerInfo.SellerId].Name,
				Rating: sellers[sellerInfo.SellerId].Rating,
				Price:  sellerInfo.Price,
			}
			itemSellersResponse = append(itemSellersResponse, sellerResponse)
		}

		itemResponse := response.ItemResponse{
			Id:              item.Id,
			Name:            item.Name,
			Description:     item.Description,
			DefaultSellerId: item.DefaultSellerId,
			SellerInfo:      itemSellersResponse,
		}

		itemsArray = append(itemsArray, itemResponse)
	}

	itemsResponse := &response.ItemsResponse{
		Items: itemsArray,
	}

	return itemsResponse, nil
}
