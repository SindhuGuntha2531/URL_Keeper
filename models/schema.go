package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//defining struct for url keeper
type Link struct {
	gorm.Model

	Name        string
	Url         string `gorm:"unique_index"`
	Category    string
	Environment string
}
