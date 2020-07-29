package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AnswerOption.Model", func() {
	var (
		course       models.Course
		level        models.Level
		targetGroup  models.TargetGroup
		target       models.Target
		quiz         models.Quiz
		quizQuestion models.QuizQuestion
		answerOption models.AnswerOption
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)
		quiz = CreateQuiz(target)
		quizQuestion = CreateQuizQuestion(quiz)
		answerOption = CreateAnswerOption(quizQuestion)

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

		qq := models.QuizQuestion{}
		qq.Delete()

		ao := models.AnswerOption{}
		ao.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new answerOption", func() {
				createdAO := models.AnswerOption{}
				createdAO.SetID(answerOption.GetID())
				createdAO.FetchByID()
				Expect(createdAO.Value).To(Equal(answerOption.Value))
			})

			It("should update an existing answer option", func() {
				answerOption.Value = "Changed Quiz Question Question"
				answerOption.UpdateOne()
				ao := models.AnswerOption{}
				ao.SetID(answerOption.GetID())
				ao.FetchByID()
				Expect(answerOption.Value).To(Equal(ao.Value))
			})

			It("should soft delete a answer option record", func() {
				answerOption.SoftDelete()
				ao := models.AnswerOption{}
				ao.SetID(answerOption.GetID())
				ao.FetchByID()
				Expect(ao.Value).To(Equal(""))
				Expect(ao.QuizQuestionID).To(Equal(""))
			})
		})
	})

	Describe("Relationship tests", func() {
		Context("Quiz Based", func() {
			It("should get the quiz question", func() {
				answerOption.GetQuestion()
				Expect(answerOption.QuizQuestion.GetID()).To(Equal(quizQuestion.GetID()))
			})
		})
	})
})
