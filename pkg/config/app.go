package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "sayam:Sayam2710@@tcp(127.0.0.1:3306)/first?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
