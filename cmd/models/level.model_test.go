package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-lms-of-pupilfirst/cmd/models"
)

var _ = Describe("Level.Model", func() {
	var (
		user   models.User
		course models.Course
		level  models.Level
	)

	BeforeEach(func() {
		// init variables
		user = CreateUser()
		course = CreateCourse()
		level = CreateLevel(course)

		if err := user.Create(); err != nil {
			Fail("Couldn't create user")
		}
	})

	AfterEach(func() {
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		l := models.Level{}
		l.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new level", func() {
				createdLevel := models.Level{}
				createdLevel.SetID(level.GetID())
				createdLevel.FetchByID()
				Expect(createdLevel.CourseID).To(Equal(course.GetID()))
			})

			It("should update an existing level", func() {
				level.Name = "Changed Level Name"
				level.UpdateOne()
				l := models.Level{}
				l.SetID(level.GetID())
				l.FetchByID()
				Expect(level.Name).To(Equal(l.Name))
			})

			It("should softdelete an exsting level", func() {
				level.SoftDelete()
				l := models.Level{}
				l.SetID(level.GetID())
				l.FetchByID()
				Expect(l.Name).To(Equal(""))
				Expect(l.Number).To(Equal(0))
			})
		})
	})

	Describe("Basic Relationship Tests", func() {
		Context("Course Based", func() {
			It("should get the course", func() {
				level.GetCourse()
				Expect(level.Course.Name).To(Equal(course.Name))
			})

			It("should get target groups of the level", func() {
				CreateTargetGroup(level)
				level.GetTargetGroups()
				Expect(len(level.TargetGroups)).To(Equal(1))
			})
		})
	})
})
