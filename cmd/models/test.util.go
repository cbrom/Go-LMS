package models

import (
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/database"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ConnectToTestDatabase connects to a test db (defined in .env)
func ConnectToTestDatabase() *gorm.DB {
	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database configuration %+v", err)
	}
	db, err := database.Initialize(dbConfig.Storage)
	SetRepoDB(db)
	return db
}

// CreateUser creates a user mock for testing
func CreateUser() User {
	password := "password"
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := password + passwordSalt
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

	user := User{
		Email:        "test1@gmail.com",
		Role:         2,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		Name:         "Test Name",
		About:        "About User",
	}
	return user
}

// CreateCourse creates a course mock for testing
func CreateCourse() Course {
	course := Course{
		Name:                "Test Course",
		Description:         "This is a test course",
		EnableLeadboard:     true,
		PublicSignup:        true,
		Featured:            true,
		About:               "This is about text, described after browsing",
		ProgressionBehavior: "progress",
		ProgressionLimit:    2,
	}

	return course
}

// CreateLevel creates a level mock for testing
func CreateLevel(course Course) Level {
	level := Level{
		Name:        "Test Level",
		CourseID:    course.GetID(),
		Description: "Level description",
		Number:      1,
	}

	return level
}

// AssignAuthor assignes a user to a course
func AssignAuthor(user User, course Course) CourseAuthor {
	courseAuthor := CourseAuthor{
		UserID:   user.GetID(),
		CourseID: course.GetID(),
	}
	return courseAuthor
}

// CreateStudentCourse creates a student and assigns it a course for testing
func CreateStudentCourse(user User, course Course) StudentCourse {
	studentCourse := StudentCourse{
		UserID:   user.GetID(),
		CourseID: course.GetID(),
	}
	return studentCourse
}
