package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kyc_urls/config"
	"kyc_urls/routes"
)

func main() {

	router := gin.New() //creating the router
	config.Connect()    //connecting to the database

	//different api end points to perform CRUD operations
	routes.GetUrlRoutes(router)
	routes.GetByNameRoutes(router)
	routes.GetByCatRoutes(router)
	routes.GetByEnvRoutes(router)
	routes.CreateRoutes(router)
	routes.DeleteRoutes(router)

	router.Run(":8080") //port

}
