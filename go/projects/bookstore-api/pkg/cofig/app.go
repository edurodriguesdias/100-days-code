package cofig

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:root/bookstore-db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = database
}

func GetDb() *gorm.DB {
	return db
}
