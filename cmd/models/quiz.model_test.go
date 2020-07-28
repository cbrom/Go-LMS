package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Quiz.Model", func() {
	var (
		course      models.Course
		level       models.Level
		targetGroup models.TargetGroup
		target      models.Target
		quiz        models.Quiz
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)
		quiz = CreateQuiz(target)

	})

	AfterEach(func() {
		c := models.Course{}
		c.Delete()

		l := models.Level{}
		l.Delete()

		tg := models.TargetGroup{}
		tg.Delete()

		t := models.Target{}
		t.Delete()

		q := models.Quiz{}
		q.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new quiz", func() {
				createdQuiz := models.Quiz{}
				createdQuiz.SetID(quiz.GetID())
				createdQuiz.FetchByID()
				Expect(createdQuiz.Title).To(Equal(quiz.Title))
			})

			It("should update an existing quiz", func() {
				quiz.Title = "Changed Quiz Title"
				quiz.UpdateOne()
				t := models.Quiz{}
				t.SetID(quiz.GetID())
				t.FetchByID()
				Expect(quiz.Title).To(Equal(t.Title))
			})

			It("should soft delete a quiz record", func() {
				quiz.SoftDelete()
				q := models.Quiz{}
				q.SetID(quiz.GetID())
				q.FetchByID()
				Expect(q.Title).To(Equal(""))
				Expect(q.TargetID).To(Equal(""))
			})
		})
	})

	Describe("Relationship tests", func() {

		var (
			quizQuestion models.QuizQuestion
		)

		BeforeEach(func() {
			quizQuestion = CreateQuizQuestion(quiz)
		})

		AfterEach(func() {
			qq := models.QuizQuestion{}
			qq.Delete()

			ao := models.AnswerOption{}
			ao.Delete()
		})
		Context("Target Based", func() {
			It("should get the target", func() {
				quiz.GetTarget()
				Expect(quiz.Target.GetID()).To(Equal(target.GetID()))
			})
		})

		Context("Quiz Questions Based", func() {
			It("should get quiz questions", func() {
				quiz.GetQuizQuestions()
				Expect(len(quiz.QuizQuestions)).To(Equal(1))
				Expect(quiz.QuizQuestions[0].GetID()).To(Equal(quizQuestion.GetID()))
			})
		})
	})
})
