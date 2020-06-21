package database

import "github.com/fun-go/ecom-service/models"

var itemsData map[string]*models.Item
var sellersData map[string]*models.Seller

func InitItemsCache(){
	itemsData = make(map[string]*models.Item)
	sellersData = make(map[string]*models.Seller)
}


func GetAllItemsInStore() []*models.Item {

	items := make([]*models.Item, len(itemsData))
	for _, item := range itemsData {
		items = append(items, item)
	}

	return items
}

func GetAllSellers() map[string]*models.Seller {
	return sellersData
}