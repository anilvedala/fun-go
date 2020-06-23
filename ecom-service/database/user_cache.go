package database

import (
	"errors"
	"github.com/fun-go/ecom-service/models"
)

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

func SaveUserDetails(user models.User) error {

	if usersData == nil {
		return errors.New("user data store is not yet initialised")
	}

	usersData[user.Id] = user

	return nil
}

