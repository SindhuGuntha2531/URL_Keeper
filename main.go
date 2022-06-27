package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"

	"os"
)

type Link struct {
	gorm.Model

	Name        string
	Url         string `gorm:"unique_index"`
	Category    string
	Environment string
}

var db *gorm.DB
var err error

func main() {
	//loading environment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully connected to database!!")
	}
	defer db.Close()

	//Make migrations to the database if they have not already been created
	db.AutoMigrate(&Link{})

	router := gin.New() //router

	//different routes

	router.GET("/", func(c *gin.Context) {
		c.String(200, "HELLO PEOPLE!")
	})

	router.GET("/links", func(c *gin.Context) {
		var urls []Link
		db.Find(&urls)
		c.JSON(200, &urls)
	})

	router.GET("/:id", func(c *gin.Context) {
		var urls Link
		fmt.Println(c.Param("id"))
		row := db.Where("Name=?", c.Param("id")).First(&urls)
		if row.RowsAffected == 0 {
			c.JSON(200, "No matched name found")
		} else if row.Error != nil {
			c.JSON(404, row.Error)
		} else {
			c.JSON(200, &urls)
		}

	})

	router.GET("/cat/:id", func(c *gin.Context) {
		var urls []Link
		row := db.Where("Category=?", c.Param("id")).Find(&urls)
		err := row.Error
		if err != nil {
			c.JSON(404, err)
		} else if row.RowsAffected == 0 {
			c.JSON(200, "No matched category found")
		} else {
			c.JSON(200, &urls)
		}

	})

	router.GET("/env/:id", func(c *gin.Context) {
		var urls []Link
		row := db.Where("Environment=?", c.Param("id")).Find(&urls)
		err := row.Error
		if err != nil {
			c.JSON(404, err)
		} else if row.RowsAffected == 0 {
			c.JSON(200, "No matched environment found")
		} else {
			c.JSON(200, &urls)
		}

	})

	router.POST("/links", func(c *gin.Context) {
		var url Link
		c.BindJSON(&url)
		db.Create(&url)
		c.JSON(200, &url)

	})

	router.DELETE("/:id", func(c *gin.Context) {
		var url Link
		db.Where("Name=?", c.Param("id")).Delete(&url)
		c.JSON(200, "Successfully deleted")
	})

	router.Run(":8080")

}
