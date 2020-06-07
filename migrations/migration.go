package migrations

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/jinzhu/gorm"
)

// Migrate migrates gorm models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
