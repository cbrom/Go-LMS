package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"go-lms-of-pupilfirst/cmd/models"

	"github.com/pborman/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var _ = Describe("User.Model", func() {
	var (
		user   models.User
		course models.Course
	)

	BeforeEach(func() {

		password := "password"
		passwordSalt := uuid.NewRandom().String()
		saltedPassword := password + passwordSalt
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)

		user = models.User{
			Email:        "test1@gmail.com",
			Role:         2,
			PasswordSalt: passwordSalt,
			PasswordHash: passwordHash,
			Name:         "Test Name",
			About:        "About user",
		}

		if err := user.Create(); err != nil {
			Fail("Couldn't create user")
		}

	})

	AfterEach(func() {
		// drop all users
		u := models.User{}
		u.Delete()
	})

	Describe("Basic Crud Tests", func() {
		Context("CRUD basics", func() {
			It("should contain the newly created user", func() {
				createdUser := &models.User{}
				createdUser.SetID(user.GetID())
				createdUser.FetchByID()
				Expect(createdUser.Name).To(Equal(user.Name))

			})

			It("should fetch a user by ID", func() {
				u := models.User{}
				u.SetID(user.GetID())
				u.FetchByID()
				Expect(user.Email).To(Equal(u.Email))
				Expect(user.PasswordHash).To(Equal(u.PasswordHash))
			})

			It("should update an existing user", func() {

				user.Name = "New Name"
				user.UpdateOne()
				u := models.User{}
				u.SetID(user.GetID())
				u.FetchByID()
				Expect(user.Name).To(Equal(u.Name))
			})

			It("should softdelete an existing user", func() {
				user.SoftDelete()
				u := models.User{}
				u.SetID(user.GetID())
				u.FetchByID()
				Expect(u.Name).To(Equal(""))
				Expect(u.Email).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship tests", func() {

		var (
			courseAuthor  models.CourseAuthor
			studentCourse models.StudentCourse
		)

		BeforeEach(func() {
			course = CreateCourse()
			if err := course.Create(); err != nil {
				Fail("Couldn't create course")
			}
		})

		AfterEach(func() {
			c := models.Course{}
			c.Delete()
		})

		Context("Course relationships", func() {
			BeforeEach(func() {
				// assign author
				courseAuthor = AssignAuthor(user, course)

				if err := courseAuthor.Create(); err != nil {
					Fail("Couldn't create course author")
				}
			})

			AfterEach(func() {
				ca := models.CourseAuthor{}
				ca.Delete()
			})
			It("should get courses authored", func() {
				user.GetAuthoredCourses()
				Expect(len(user.AuthoredCourses)).To(Equal(1))
			})
		})

		Context("Student relationships", func() {
			BeforeEach(func() {
				// create a student
				studentCourse = CreateStudentCourse(user, course)

				if err := studentCourse.Create(); err != nil {
					Fail("Couldn't create student course")
				}
			})

			AfterEach(func() {
				sc := models.StudentCourse{}
				sc.Delete()
			})
			It("should get student's course list", func() {
				user.GetCourses()
				Expect(len(user.Courses)).To(Equal(1))
			})
		})
	})
})
