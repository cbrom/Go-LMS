package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Certificate.Model", func() {
	var (
		user         models.User
		course       models.Course
		courseAuthor models.CourseAuthor
		certificate  models.Certificate
	)

	BeforeEach(func() {
		// init data
		user = CreateUser()
		course = CreateCourse()
		courseAuthor = AssignAuthor(user, course)
		certificate = CreateCertificate(courseAuthor)
	})

	AfterEach(func() {
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		ca := models.CourseAuthor{}
		ca.Delete()

		cer := models.Certificate{}
		cer.Delete()
	})
	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should fetcha certificate ID", func() {
				c := models.Certificate{}
				c.SetID(certificate.GetID())
				c.FetchByID()
				Expect(certificate.CourseID).To(Equal(c.CourseID))
				Expect(certificate.CourseAuthorID).To(Equal(c.CourseAuthorID))
			})

			It("should update an existing certificate", func() {
				certificate.QRCorner = "Changed QRCorner"
				certificate.UpdateOne()
				c := models.Certificate{}
				c.SetID(certificate.GetID())
				c.FetchByID()
				Expect(c.QRCorner).To(Equal(certificate.QRCorner))
			})

			It("should softdelete an existing certificate", func() {
				certificate.SoftDelete()
				c := models.Certificate{}
				c.SetID(certificate.GetID())
				c.FetchByID()
				Expect(c.QRCorner).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship Tests", func() {
		Context("Course Based", func() {
			It("should get the course", func() {
				certificate.GetCourse()
				Expect(certificate.Course.GetID()).To(Equal(course.GetID()))
			})
		})

		Context("Issuer Based", func() {
			It("should get the issuer", func() {
				certificate.GetIssuer()
				Expect(certificate.Issuer.GetID()).To(Equal(courseAuthor.GetID()))
			})
		})
	})
})
