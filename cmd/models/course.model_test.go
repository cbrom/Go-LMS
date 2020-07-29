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

	Describe("Basic Relationship Tests", func() {
		var (
			level         models.Level
			studentCourse models.StudentCourse
			courseAuthor  models.CourseAuthor
			certificate   models.Certificate
		)

		BeforeEach(func() {
			level = CreateLevel(course)
			studentCourse = CreateStudentCourse(user, course)
			courseAuthor = AssignAuthor(user, course)
			certificate = CreateCertificate(courseAuthor)
		})

		AfterEach(func() {
			l := models.Level{}
			l.Delete()

			s := models.StudentCourse{}
			s.Delete()

			ca := models.CourseAuthor{}
			ca.Delete()
		})

		Context("Level Based", func() {
			It("should get course levels", func() {
				course.GetLevels()
				Expect(len(course.Levels)).To(Equal(1))
				Expect(course.Levels[0].GetID()).To(Equal(level.GetID()))
			})
		})

		Context("Author based", func() {
			It("should get course authors", func() {
				course.GetAuthors()
				Expect(len(course.Authors)).To(Equal(1))
				Expect(course.Authors[0].GetID()).To(Equal(courseAuthor.GetID()))
			})
		})

		Context("Student based", func() {
			It("should course students", func() {
				course.GetStudents()
				Expect(len(course.Students)).To(Equal(1))
				Expect(course.Students[0].GetID()).To(Equal(studentCourse.GetID()))
			})
		})

		Context("Certificate based", func() {
			It("should get course certificates", func() {
				course.GetCertificates()
				Expect(len(course.Certificates)).To(Equal(1))
				Expect(course.Certificates[0].GetID()).To(Equal(certificate.GetID()))
			})
		})
	})

})
