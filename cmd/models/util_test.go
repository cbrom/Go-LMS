package models_test

import (
	"encoding/json"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/database"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
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
	certificate := models.Certificate{
		CourseID:       courseAuthor.Course.GetID(),
		CourseAuthorID: courseAuthor.GetID(),
		QRCorner:       "QR Corner",
		QRScale:        2,
		Margin:         2,
		NameOffsetTop:  2,
		FontSize:       2,
		Message:        "Certificate message",
		Active:         true,
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

	if err := studentCourse.Create(); err != nil {
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

func CreateTarget(targetGroup models.TargetGroup) models.Target {
	target := models.Target{
		Role:                   "Target Test role",
		Title:                  "Target Test Title",
		Description:            "Target Test Description",
		CompletionInstructions: "",
		ResourceURL:            "Target Test resource url",
		TargetGroupID:          targetGroup.GetID(),
		SortIndex:              1,
		LinkToComplete:         "Target Test no link",
		Resubmittable:          true,
	}

	if err := target.Create(); err != nil {
		Fail("Couldn't create target")
	}

	return target
}

func CreateTargetVersion(target models.Target) models.TargetVersion {
	targetVersion := models.TargetVersion{
		TargetID:    target.GetID(),
		VersionName: "Target Version Test Name",
	}
	if err := targetVersion.Create(); err != nil {
		Fail("Couldn't create target")
	}

	return targetVersion
}

func CreateContentBlock(targetVersion models.TargetVersion) models.ContentBlock {
	value := struct {
		Key   string
		Value string
	}{Key: "autorefid", Value: "100"}

	returned, _ := json.Marshal(value)
	contentBlock := models.ContentBlock{
		BlockType:       "text",
		Content:         postgres.Jsonb{returned},
		SortIndex:       1,
		TargetVersionID: targetVersion.GetID(),
	}
	if err := contentBlock.Create(); err != nil {
		Fail("Couldn't create content block")
	}

	return contentBlock
}

func CreateQuiz(target models.Target) models.Quiz {
	quiz := models.Quiz{
		Title:    "Quiz Test Title",
		TargetID: target.GetID(),
	}
	if err := quiz.Create(); err != nil {
		Fail("Couldn't create quiz")
	}

	return quiz
}

func CreateQuizQuestion(quiz models.Quiz) models.QuizQuestion {
	quizQuestion := models.QuizQuestion{
		QuizID:      quiz.GetID(),
		Question:    "This is a Test question",
		Description: "this is a question to help students get good at math",
	}
	if err := quizQuestion.Create(); err != nil {
		Fail("Couldn't create quizQuestion")
	}

	return quizQuestion
}

func CreateAnswerOption(quizQuestion models.QuizQuestion) models.AnswerOption {
	answerOption := models.AnswerOption{
		QuizQuestionID: quizQuestion.GetID(),
		Value:          "This is a Test answer",
		Hint:           "This is a Test hint",
	}

	if err := answerOption.Create(); err != nil {
		Fail("Couldn't create answer option")
	}

	return answerOption
}
