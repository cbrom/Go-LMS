package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"

	"go-lms-of-pupilfirst/cmd/models"
)

var _ = Describe("CourseAuthor.Model", func() {
	var (
		user         models.User
		course       models.Course
		courseAuthor models.CourseAuthor
	)

	BeforeEach(func() {
		// create user
		user = CreateUser()
		// create course
		course = CreateCourse()

		// assign author
		courseAuthor = AssignAuthor(user, course)
	})

	AfterEach(func() {
		// drop data
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		ca := models.CourseAuthor{}
		ca.Delete()
	})

	Describe("Basic Crud Tests", func() {
		Context("CRUD basics", func() {
			It("should create a new course author", func() {
				createdCA := models.CourseAuthor{}
				createdCA.SetID(courseAuthor.GetID())
				createdCA.FetchByID()
				Expect(createdCA.CourseID).To(Equal(courseAuthor.CourseID))
			})

			It("should fetcha course author by ID", func() {
				ca := models.CourseAuthor{}
				ca.SetID(courseAuthor.GetID())
				ca.FetchByID()
				Expect(courseAuthor.CourseID).To(Equal(ca.CourseID))
				Expect(courseAuthor.UserID).To(Equal(ca.UserID))
			})

			It("should update an existing course author", func() {
				newID := uuid.NewRandom().String()
				courseAuthor.UserID = newID
				courseAuthor.UpdateOne()
				ca := models.CourseAuthor{}
				ca.SetID(courseAuthor.GetID())
				ca.FetchByID()
				Expect(ca.UserID).To(Equal(newID))
			})

			It("should soft delete a record", func() {
				courseAuthor.SoftDelete()
				ca := models.CourseAuthor{}
				ca.SetID(courseAuthor.GetID())
				ca.FetchByID()
				Expect(ca.CourseID).To(Equal(""))
				Expect(ca.UserID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship Tests", func() {
		Context("Course Based", func() {
			It("should get the course", func() {
				courseAuthor.GetCourse()
				Expect(courseAuthor.Course.Name).To(Equal(course.Name))
			})
		})

		Context("Author based", func() {
			It("should get the author", func() {
				courseAuthor.GetUser()
				Expect(courseAuthor.User.Name).To(Equal(user.Name))
			})
		})

		Context("Certificate based", func() {
			BeforeEach(func() {
				CreateCertificate(courseAuthor)
			})
			AfterEach(func() {
				cer := models.Certificate{}
				cer.Delete()
			})
			It("should get certificates of the level", func() {
				courseAuthor.GetCertificates()
				Expect(len(courseAuthor.Certificates)).To(Equal(1))
			})
		})
	})
})
