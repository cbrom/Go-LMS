package models

import (
	"testing"
	"strconv"
	"time"
	"github.com/jinzhu/gorm"
	"go-lms-of-pupilfirst/configs"
	"golang.org/x/crypto/bcrypt"
	"go-lms-of-pupilfirst/pkg/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/pborman/uuid"
)

const (
	success = "\u2713"
	failed = "\u2717"
)

func connectToDB(t *testing.T) (*gorm.DB, error) {
	if err := envconfig.Process("go-lms-api", &configs.CFG); err != nil {
		t.Errorf("main : Error Parsing Config file: %+v", err)
	}

	dbConfig, err := configs.LoadConfig()
	if err != nil {
		t.Errorf("main : Error loading database %+v", err)
	}
	t.Logf("%+v", dbConfig)
	db, err := database.Initialize(dbConfig.Storage)
	return db, err
}
func purge(db *gorm.DB) {
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&AnswerOption{})
	db.DropTableIfExists(&Certificate{})
	db.DropTableIfExists(&IssuedCertificate{})
	db.DropTableIfExists(&ContentBlock{})
	db.DropTableIfExists(&CourseAuthor{})
	db.DropTableIfExists(&Course{})
	db.DropTableIfExists(&EvaluationCriteria{})
	db.DropTableIfExists(&Level{})
	db.DropTableIfExists(&QuizQuestion{})
	db.DropTableIfExists(&QuizUserAnswer{})
	db.DropTableIfExists(&Quiz{})
	db.DropTableIfExists(&Target{})
	db.DropTableIfExists(&TargetVersion{})
	db.DropTableIfExists(&TargetVersion{})
	db.DropTableIfExists(&TargetGroup{})
	db.DropTableIfExists(&StudentCourse{})
}
func migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&AnswerOption{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Certificate{})
	db.AutoMigrate(&IssuedCertificate{})
	db.AutoMigrate(&ContentBlock{})
	db.AutoMigrate(&CourseAuthor{})
	db.AutoMigrate(&EvaluationCriteria{})
	db.AutoMigrate(&Level{})
	db.AutoMigrate(&QuizQuestion{})
	db.AutoMigrate(&QuizUserAnswer{})
	db.AutoMigrate(&Quiz{})
	db.AutoMigrate(&Target{})
	db.AutoMigrate(&TargetVersion{})
	db.AutoMigrate(&TargetVersion{})
	db.AutoMigrate(&TargetGroup{})
	db.AutoMigrate(&StudentCourse{})
}

func encrypt(password string) (string, []byte){
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return passwordSalt, passwordHash
	}
	return "", []byte{}
}
func createUser(name string, email string, password string, timeZone time.Time) *User {
	passwordSalt, passwordHash := encrypt(password)
	usr := &User{
		Name:         name,
		Email:        email,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		TimeZone:     &timeZone,
	}	
	usr.Create()
	return usr
}
func createCourseAuthor(course *Course, user *User) *CourseAuthor{
	courseAuthor := &CourseAuthor{
		CourseID:       course.GetID(),
		UserID:         user.GetID(),
	}
	courseAuthor.Create()
	return courseAuthor
}
func createCourse(name string, endsAt time.Time, description string, enableLeadboard bool, publicSignup bool, featured bool, about string, progressionBehavior string, progressionLimit int) *Course{
	course := &Course{
		Name: 				 name,
		EndsAt: 			 &endsAt,
		Description:         description,
		EnableLeadboard:     enableLeadboard,
		PublicSignup:        publicSignup,
		Featured:            featured,
		About:               about,
		ProgressionBehavior: progressionBehavior,
		ProgressionLimit:    progressionLimit,
	}
	course.Create()
	return course
}
func createCertificate(course *Course, issuer *CourseAuthor, qrCorner string, qrScale int, margin int, nameOffsetTop int, fontSize int, message string, active bool) *Certificate {
	certificate := &Certificate{
		CourseID:	   course.GetID(),
		IssuerID:	   issuer.GetID(),
		QRCorner:      qrCorner,
		QRScale:       qrScale,
		Margin:        margin,
		NameOffsetTop: nameOffsetTop,
		FontSize:      fontSize,
		Message:       message,
		Active:        active,

	}
	certificate.Create()
	return certificate
}

func TestGetCertificates(t *testing.T) {
	db, err := connectToDB(t)
	if err == nil {
		defer db.Close()
		SetRepoDB(db)
		purge(db)
		migrate(db)
	} else {
		t.Errorf("Database is not setup")
	}
	n := 10
	currentTime := time.Now()
	user := createUser("cbrom", "cbrom@gmail.com", "password", currentTime)
	course := createCourse("course", time.Now(), "1", false, false, false, "1", uuid.NewRandom().String(), 0)
	courseAuthor := createCourseAuthor(course, user)
	t.Logf("coureAuthor %v", courseAuthor)
	certificates := make(map[string]*Certificate, n)
	for i := 1; i<=n ; i++ {
		strI := strconv.Itoa(i)
		certificate := createCertificate(course, courseAuthor, strI, i, 0, 0, 0, uuid.NewRandom().String(), true)
		certificates[certificate.GetID()] = certificate 
	}
	userFromDB := &User{}
	userFromDB.SetID(user.GetID())
	userFromDB.FetchByID()	
	courseFromDB := &Course{}
	courseFromDB.SetID(course.GetID())
	courseFromDB.FetchByID()
	err = courseFromDB.GetCertificates()
	for _, c := range courseFromDB.Certificates {
		// Time created can be off by a few milliseconds
		c.Base.CreatedAt = c.Base.CreatedAt.Truncate(time.Second)
		c.Base.UpdatedAt = c.Base.UpdatedAt.Truncate(time.Second)
		val, ok := certificates[c.GetID()]
		t.Logf("Given that we created certificates with course and courseAuthor")
		{
			t.Logf("\tWhen we call courseFromDB.GetCertificates()")
			if ok {
				val.Base.CreatedAt = val.Base.CreatedAt.Truncate(time.Second)
				val.Base.UpdatedAt = val.Base.UpdatedAt.Truncate(time.Second)
				if (!c.Base.CreatedAt.Equal(val.Base.CreatedAt) || !c.Base.UpdatedAt.Equal(val.Base.UpdatedAt) || (c.GetID() != val.GetID()) || (c.CourseID != val.CourseID) || (c.QRCorner != val.QRCorner) || (c.QRScale != val.QRScale) || (c.Margin != val.Margin) || (c.NameOffsetTop != val.NameOffsetTop) || (c.FontSize != val.FontSize) || (c.Message != val.Message) || (c.Active != val.Active)) {
					t.Errorf("\t%s\tThe certificate from database with id %s != certificate created. \n%v \n!= \n%v", failed, c.GetID(), *c, *val)
				}
				t.Logf("\t%s\tCertificates created should be the same as the ones retrieved from database", success)
			} else {
				t.Errorf("\t%s\tCertificate from database with id %s != certificate is not expected", failed, c.GetID())
			}

		}
	}
} 
