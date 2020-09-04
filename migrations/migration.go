package migrations

import (
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/jinzhu/gorm"
)

// Migrate migrates gorm models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.CourseAuthor{})
	db.AutoMigrate(&models.Certificate{})
	db.AutoMigrate(&models.IssuedCertificate{})
	db.AutoMigrate(&models.StudentCourse{})
	db.AutoMigrate(&models.EvaluationCriteria{})
	db.AutoMigrate(&models.Level{})
	db.AutoMigrate(&models.TargetGroup{})
	db.AutoMigrate(&models.Target{})
	db.AutoMigrate(&models.TargetVersion{})
	db.AutoMigrate(&models.ContentBlock{})
	db.AutoMigrate(&models.Quiz{})
	db.AutoMigrate(&models.QuizQuestion{})
	db.AutoMigrate(&models.AnswerOption{})
	db.AutoMigrate(&models.QuizUserAnswer{})

}
