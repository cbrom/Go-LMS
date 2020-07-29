package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-lms-of-pupilfirst/cmd/models"
)

var _ = Describe("EvaluationCriteria.Model", func() {
	var (
		course             models.Course
		evaluationCriteria models.EvaluationCriteria
	)

	BeforeEach(func() {
		// create course
		course = CreateCourse()
		evaluationCriteria = CreateEvaluationCriteria(course)
	})

	AfterEach(func() {
		// drop datas

		c := models.Course{}
		c.Delete()

		ec := models.EvaluationCriteria{}
		ec.Delete()
	})

	Describe("Basic Crud Tests", func() {
		Context("CRUD basics", func() {
			It("should create a new evaluation criteria", func() {
				createdEC := models.EvaluationCriteria{}
				createdEC.SetID(evaluationCriteria.GetID())
				createdEC.FetchByID()
				Expect(createdEC.Name).To(Equal(evaluationCriteria.Name))
			})

			It("should fetch a evaluation criteria by ID", func() {
				ec := models.EvaluationCriteria{}
				ec.SetID(evaluationCriteria.GetID())
				ec.FetchByID()
				Expect(evaluationCriteria.PassGrade).To(Equal(ec.PassGrade))
				Expect(evaluationCriteria.Name).To(Equal(ec.Name))
			})

			It("should update an existing evaluation criteria", func() {
				evaluationCriteria.Name = "Changed Evaluation Criteria Name"
				evaluationCriteria.UpdateOne()
				ec := models.EvaluationCriteria{}
				ec.SetID(evaluationCriteria.GetID())
				ec.FetchByID()
				Expect(ec.Name).To(Equal(evaluationCriteria.Name))
			})

			It("should softdelete an existing user", func() {
				evaluationCriteria.SoftDelete()
				ec := models.EvaluationCriteria{}
				ec.SetID(course.GetID())
				ec.FetchByID()
				Expect(ec.Name).To(Equal(""))
				Expect(ec.PassGrade).To(Equal(uint(0)))
			})
		})
	})

	Describe("Relationship Tests", func() {
		Context("Course Based", func() {
			It("should get the course", func() {
				evaluationCriteria.GetCourse()
				Expect(evaluationCriteria.Course.GetID()).To(Equal(course.GetID()))
			})
		})
	})

})
