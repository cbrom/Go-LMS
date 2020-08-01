package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"

	"go-lms-of-pupilfirst/cmd/models"
)

var _ = Describe("IssuedCertificate.Model", func() {
	var (
		user              models.User
		course            models.Course
		courseAuthor      models.CourseAuthor
		certificate       models.Certificate
		issuedCertificate models.IssuedCertificate
	)

	BeforeEach(func() {
		// init data
		user = CreateUser()
		course = CreateCourse()
		courseAuthor = AssignAuthor(user, course)
		certificate = CreateCertificate(courseAuthor)
		issuedCertificate = CreateIssuedCertificate(certificate, user)
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

		ic := models.IssuedCertificate{}
		ic.Delete()
	})
	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should fetch an issued certificate", func() {
				ic := models.IssuedCertificate{}
				ic.SetID(issuedCertificate.GetID())
				ic.FetchByID()
				Expect(issuedCertificate.SerialNumber).To(Equal(ic.SerialNumber))
			})

			It("should update an existing issued certificate", func() {
				issuedCertificate.SerialNumber = uuid.NewRandom().String()
				issuedCertificate.UpdateOne()
				ic := models.IssuedCertificate{}
				ic.SetID(issuedCertificate.GetID())
				ic.FetchByID()
				Expect(ic.SerialNumber).To(Equal(issuedCertificate.SerialNumber))
			})

			It("should softdelete an existing issued certificate", func() {
				issuedCertificate.SoftDelete()
				ic := models.IssuedCertificate{}
				ic.SetID(issuedCertificate.GetID())
				ic.FetchByID()
				Expect(ic.SerialNumber).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship test", func() {
		Context("Certificate Based", func() {
			It("should get the certificate", func() {
				issuedCertificate.GetCertificate()
				Expect(issuedCertificate.Certificate.GetID()).To(Equal(certificate.GetID()))
			})
		})
	})
})
