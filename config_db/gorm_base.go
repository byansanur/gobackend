package config_db

import "github.com/jinzhu/gorm"

type InDB struct {
	DB *gorm.DB
}
