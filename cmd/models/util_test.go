package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/database"
	"log"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("Util", func() {

})

// ConnectToTestDatabase connects to a test db (defined in .env)
func ConnectToTestDatabase() *gorm.DB {
	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database configuration %+v", err)
	}
	db, err := database.Initialize(dbConfig.Storage)
	models.SetRepoDB(db)
	return db
}

// CreateUser creates a user mock for testing
func CreateUser() models.User {
	password := "password"
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := password + passwordSalt
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

	user := models.User{
		Email:        "test1@gmail.com",
		Role:         2,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		Name:         "Test Name",
		About:        "About User",
	}
	if err := user.Create(); err != nil {
		Fail("Couldn't create user")
	}
	return user
}

// CreateCourse creates a course mock for testing
func CreateCourse() models.Course {
	course := models.Course{
		Name:                "Test Course",
		Description:         "This is a test course",
		EnableLeadboard:     true,
		PublicSignup:        true,
		Featured:            true,
		About:               "This is about text, described after browsing",
		ProgressionBehavior: "progress",
		ProgressionLimit:    2,
	}

	if err := course.Create(); err != nil {
		Fail("Couldn't create course")
	}

	return course
}

// CreateLevel creates a level mock for testing
func CreateLevel(course models.Course) models.Level {
	level := models.Level{
		Name:        "Test Level",
		CourseID:    course.GetID(),
		Description: "Level description",
		Number:      1,
	}

	if err := level.Create(); err != nil {
		Fail("Couldn't create level")
	}

	return level
}

// AssignAuthor assignes a user to a course
func AssignAuthor(user models.User, course models.Course) models.CourseAuthor {
	courseAuthor := models.CourseAuthor{
		UserID:   user.GetID(),
		CourseID: course.GetID(),
	}

	if err := courseAuthor.Create(); err != nil {
		Fail("Couldn't create course author")
	}
	return courseAuthor
}

// CreateCertificate creates a certificate mock for testing
func CreateCertificate(courseAuthor models.CourseAuthor) models.Certificate {
	courseAuthor.GetCourse()
	courseAuthor.GetUser()
	certificate := models.Certificate{
		CourseID:      courseAuthor.Course.GetID(),
		IssuerID:      courseAuthor.User.GetID(),
		QRCorner:      "QR Corner",
		QRScale:       2,
		Margin:        2,
		NameOffsetTop: 2,
		FontSize:      2,
		Message:       "Certificate message",
		Active:        true,
	}

	if err := certificate.Create(); err != nil {
		Fail("Couldn't create certificate")
	}
	return certificate
}

// CreateStudentCourse creates a student and assigns it a course for testing
func CreateStudentCourse(user models.User, course models.Course) models.StudentCourse {
	studentCourse := models.StudentCourse{
		UserID:   user.GetID(),
		CourseID: course.GetID(),
	}

	if err := user.Create(); err != nil {
		Fail("Couldn't create student course")
	}

	return studentCourse
}

// CreateTargetGroup creates a level target group mock for testing
func CreateTargetGroup(level models.Level) models.TargetGroup {
	targetGroup := models.TargetGroup{
		Name:        "Test target group",
		Description: "Test target description",
		SortIndex:   1,
		Milestone:   true,
		LevelID:     level.GetID(),
	}

	if err := targetGroup.Create(); err != nil {
		Fail("Couldn't create target group")
	}

	return targetGroup
}
