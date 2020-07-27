package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
)

var _ = Describe("StudentCourse.Model", func() {
	var (
		user          models.User
		course        models.Course
		studentCourse models.StudentCourse
	)

	BeforeEach(func() {
		// create user
		user = CreateUser()
		// create course
		course = CreateCourse()

		// create student course
		studentCourse = CreateStudentCourse(user, course)
	})

	AfterEach(func() {
		// drop data
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		ca := models.StudentCourse{}
		ca.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should fetcha student course by ID", func() {
				sc := models.StudentCourse{}
				sc.SetID(studentCourse.GetID())
				sc.FetchByID()
				Expect(studentCourse.CourseID).To(Equal(sc.CourseID))
				Expect(studentCourse.UserID).To(Equal(sc.UserID))
			})

			It("should update an existing student course", func() {
				newID := uuid.NewRandom().String()
				studentCourse.UserID = newID
				studentCourse.UpdateOne()
				ca := models.StudentCourse{}
				ca.SetID(studentCourse.GetID())
				ca.FetchByID()
				Expect(ca.UserID).To(Equal(newID))
			})

			It("should soft delete a student course record", func() {
				studentCourse.SoftDelete()
				ca := models.StudentCourse{}
				ca.SetID(studentCourse.GetID())
				ca.FetchByID()
				Expect(ca.CourseID).To(Equal(""))
				Expect(ca.UserID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship Tests", func() {
		Context("Course Based", func() {
			It("should get the course", func() {
				studentCourse.GetCourse()
				Expect(studentCourse.Course.Name).To(Equal(course.Name))
			})
		})

		Context("Author based", func() {
			It("should get the student", func() {
				studentCourse.GetUser()
				Expect(studentCourse.User.Name).To(Equal(user.Name))
			})
		})
	})
})
