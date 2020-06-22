package database

import "github.com/fun-go/ecom-service/models"

var usersData  map[string]models.User

func InitUserCache(){
	usersData = make(map[string]models.User)
}


func GetUserDetails(userId string) models.User {

	user, ok := usersData[userId]
	if !ok {
		return models.User{}
	}

	return user
}

