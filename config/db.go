package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kyc_urls/models"
	"log"
	"os"
)

var Db *gorm.DB
var err error

func Connect() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	Db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully connected to database!!")
	}

	//Make migrations to the database if they have not already been created
	Db.AutoMigrate(&models.Link{})
}
