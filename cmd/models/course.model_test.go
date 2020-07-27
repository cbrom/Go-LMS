package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Course", func() {
	var (
		user   models.User
		course models.Course
	)

	BeforeEach(func() {
		// create user
		user = CreateUser()
		// create course
		course = CreateCourse()

		if err := user.Create(); err != nil {
			Fail("Couldn't create user")
		}

		if err := course.Create(); err != nil {
			Fail("Couldn't create course")
		}
	})

	AfterEach(func() {
		// drop datas
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()
	})

	Describe("Basic Crud Tests", func() {
		Context("CRUD basics", func() {
			It("should create a new user", func() {
				createdCourse := models.Course{}
				createdCourse.SetID(course.GetID())
				createdCourse.FetchByID()
				Expect(createdCourse.Name).To(Equal(course.Name))
			})

			It("should fetch a course by ID", func() {
				c := models.Course{}
				c.SetID(course.GetID())
				c.FetchByID()
				Expect(course.Name).To(Equal(c.Name))
				Expect(course.Description).To(Equal(c.Description))
			})

			It("should update an existing course", func() {
				course.Name = "Changed Name"
				course.UpdateOne()
				c := models.Course{}
				c.SetID(course.GetID())
				c.FetchByID()
				Expect(c.Name).To(Equal("Changed Name"))
			})

			It("should softdelete an existing user", func() {
				course.SoftDelete()
				c := models.Course{}
				c.SetID(course.GetID())
				c.FetchByID()
				Expect(c.Name).To(Equal(""))
				Expect(c.Description).To(Equal(""))
			})
		})
	})

})
