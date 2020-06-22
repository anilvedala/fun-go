package main

import (
	"github.com/fun-go/ecom-service/database"
	"github.com/fun-go/ecom-service/routes"
)

func main() {

	database.InitCaches()
	routes.InitRoutes()

}
