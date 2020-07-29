package models_test

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuizQuestion.Model", func() {
	var (
		course       models.Course
		level        models.Level
		targetGroup  models.TargetGroup
		target       models.Target
		quiz         models.Quiz
		quizQuestion models.QuizQuestion
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)
		quiz = CreateQuiz(target)
		quizQuestion = CreateQuizQuestion(quiz)

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
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new quiz question", func() {
				createdQuizQuestion := models.QuizQuestion{}
				createdQuizQuestion.SetID(quizQuestion.GetID())
				createdQuizQuestion.FetchByID()
				Expect(createdQuizQuestion.Question).To(Equal(quizQuestion.Question))
			})

			It("should update an existing quiz question", func() {
				quizQuestion.Question = "Changed Quiz Question Question"
				quizQuestion.UpdateOne()
				q := models.QuizQuestion{}
				q.SetID(quizQuestion.GetID())
				q.FetchByID()
				Expect(quizQuestion.Question).To(Equal(q.Question))
			})

			It("should soft delete a quiz record", func() {
				quizQuestion.SoftDelete()
				q := models.QuizQuestion{}
				q.SetID(quizQuestion.GetID())
				q.FetchByID()
				Expect(q.Question).To(Equal(""))
				Expect(q.QuizID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship tests", func() {
		var (
			answerOption models.AnswerOption
		)

		BeforeEach(func() {
			answerOption = CreateAnswerOption(quizQuestion)
			quizQuestion.CorrectAnswerID = answerOption.GetID()
			quizQuestion.UpdateOne()
		})
		AfterEach(func() {
			ao := models.AnswerOption{}
			ao.Delete()
		})
		Context("Quiz Based", func() {
			It("shoud get the quiz", func() {
				quizQuestion.GetQuiz()
				Expect(quizQuestion.Quiz.GetID()).To(Equal(quiz.GetID()))
			})
		})

		Context("Answer Based", func() {
			It("should get the answers", func() {
				quizQuestion.GetAnswerOptions()
				Expect(len(quizQuestion.Answers)).To(Equal(1))
				Expect(quizQuestion.Answers[0].GetID()).To(Equal(answerOption.GetID()))
			})

			It("should get the correct answer", func() {
				quizQuestion.GetAnswer()
				fmt.Printf("quiz question %+v\n\n", quizQuestion.Answer.GetID())
				fmt.Printf("quiz question %+v", answerOption.GetID())
				Expect(quizQuestion.Answer.GetID()).To(Equal(answerOption.GetID()))
			})
		})
	})
})
