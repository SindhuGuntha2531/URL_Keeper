package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kyc_urls/config"
	"kyc_urls/models"
)

//defining controller functions for different end points
func GetLinks(c *gin.Context) {
	var urls []models.Link
	config.Db.Find(&urls)
	c.JSON(200, &urls)
}

//get specific link based on name
func GetLinkByName(c *gin.Context) {
	var urls models.Link
	fmt.Println(c.Param("id"))
	row := config.Db.Where("Name=?", c.Param("id")).First(&urls)
	if row.RowsAffected == 0 {
		c.JSON(200, "No matched name found")
	} else if row.Error != nil {
		c.JSON(404, row.Error)
	} else {
		c.JSON(200, &urls)
	}

}

//get links based on category
func GetLinksByCat(c *gin.Context) {
	var urls []models.Link
	row := config.Db.Where("Category=?", c.Param("id")).Find(&urls)
	err := row.Error
	if err != nil {
		c.JSON(404, err)
	} else if row.RowsAffected == 0 {
		c.JSON(200, "No matched category found")
	} else {
		c.JSON(200, &urls)
	}

}

//get links based on environment
func GetLinksByEnv(c *gin.Context) {
	var urls []models.Link
	row := config.Db.Where("Environment=?", c.Param("id")).Find(&urls)
	err := row.Error
	if err != nil {
		c.JSON(404, err)
	} else if row.RowsAffected == 0 {
		c.JSON(200, "No matched environment found")
	} else {
		c.JSON(200, &urls)
	}

}

//Adding a new link
func CreateLink(c *gin.Context) {
	var url models.Link
	c.BindJSON(&url)
	config.Db.Create(&url)
	c.JSON(200, &url)

}

//deleting existing links
func DeleteLink(c *gin.Context) {
	var url models.Link
	config.Db.Where("Name=?", c.Param("id")).Delete(&url)
	c.JSON(200, "Successfully deleted")
}
