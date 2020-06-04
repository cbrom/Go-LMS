package migrations

import (
	"go-lms-of-pupilfirst/cmd/models/user"
	"go-lms-of-pupilfirst/pkg/database"
	"log"
)

// Migrate migrates gorm models
func Migrate() {
	db, err := database.Handler().Prepare()
	if err != nil {
		log.Fatalf("Unable to migrate, %+v", err)
	}
	db.AutoMigrate(&user.User{})
}
