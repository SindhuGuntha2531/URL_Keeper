package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Link struct {
	gorm.Model

	Name        string
	Url         string `gorm:"unique_index"`
	Category    string
	Environment string
}
