package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin")
func GetLinks(c *gin.Context) {
	var urls []Link
	db.Find(&urls)
	c.JSON(200, &urls)
}

func GetLinkByName(c *gin.Context) {
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

}

func GetLinksByCat(c *gin.Context) {
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

}

func GetLinksByEnv(c *gin.Context) {
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

}

func CreateLink(c *gin.Context) {
	var url Link
	c.BindJSON(&url)
	db.Create(&url)
	c.JSON(200, &url)

}

func DeleteLink(c *gin.Context) {
	var url Link
	db.Where("Name=?", c.Param("id")).Delete(&url)
	c.JSON(200, "Successfully deleted")
}
)

func GetLinks(c *gin.Context) {
	var urls []Link
	db.Find(&urls)
	c.JSON(200, &urls)
}

func GetLinkByName(c *gin.Context) {
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

}

func GetLinksByCat(c *gin.Context) {
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

}

func GetLinksByEnv(c *gin.Context) {
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

}

func CreateLink(c *gin.Context) {
	var url Link
	c.BindJSON(&url)
	db.Create(&url)
	c.JSON(200, &url)

}

func DeleteLink(c *gin.Context) {
	var url Link
	db.Where("Name=?", c.Param("id")).Delete(&url)
	c.JSON(200, "Successfully deleted")
}
