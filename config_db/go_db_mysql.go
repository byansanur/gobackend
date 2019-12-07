package config_db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
var keys = []byte("byandev")

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp/gogo?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		err.Error()
	}
	return db
}

func JwtKey() []byte {
	return keys
}
