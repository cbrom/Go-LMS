package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	handlerNotSet error = errors.New("Handler not set properly")
	handler       *gorm.DB
)

// SetRepoDB global db handler
func SetRepoDB(db *gorm.DB) {
	handler = db
}
