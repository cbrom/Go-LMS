package models_test

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/migrations"
	"go-lms-of-pupilfirst/pkg/database"

	"github.com/pborman/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var _ = Describe("User.Model", func() {
	var (
		user     models.User
		db       *gorm.DB
		userName string
	)

	BeforeEach(func() {

		dbConfig, err := configs.LoadConfig()
		if err != nil {
			log.Printf("main : Error loading database %+v", err)
		}
		log.Printf("%+v", dbConfig.Storage)
		db, err = database.Initialize(dbConfig.Storage)

		models.SetRepoDB(db)

		migrations.Migrate(db)
		// defer db.Close()

		password := "password"
		passwordSalt := uuid.NewRandom().String()
		saltedPassword := password + passwordSalt
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

		user = models.User{
			Email:        "test1@gmail.com",
			Role:         2,
			PasswordSalt: passwordSalt,
			PasswordHash: passwordHash,
			Name:         "Test Name",
			About:        "About user",
		}

	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Registering a new User", func() {
		Context("Create a new User", func() {
			It("should contain the newly created user", func() {
				if err := user.Create(); err == nil {
					createdUser := &models.User{}
					createdUser.SetID(user.GetID())
					createdUser.FetchByID()
					userName = createdUser.Name
					Expect(userName).To(Equal(user.Name))
				} else {
					if err.Error() == "pq: duplicate key value violates unique constraint \"uix_users_email\"" {
						fmt.Printf("Error %+v", err.Error())
					} else {
						Fail("Couldn't create user")
					}

				}

			})
		})
	})
})
