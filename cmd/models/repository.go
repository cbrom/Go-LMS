package models

import "github.com/jinzhu/gorm"

var handler *gorm.DB

func SetRepoDB(db *gorm.DB) {
	handler = db
}
