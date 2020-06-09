package models

import "github.com/jinzhu/gorm"

var handler *gorm.DB

// SetRepoDB global db handler
func SetRepoDB(db *gorm.DB) {
	handler = db
}
