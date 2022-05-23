package models

import (
	"./pkg/config"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	database = config.GetDb()
}
