package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Target.Model", func() {
	var (
		course      models.Course
		level       models.Level
		targetGroup models.TargetGroup
		target      models.Target
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)

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
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new target", func() {
				createdTarget := models.Target{}
				createdTarget.SetID(target.GetID())
				createdTarget.FetchByID()
				Expect(createdTarget.Title).To(Equal(target.Title))
			})

			It("should update an existing target", func() {
				target.Title = "Changed Target Title"
				target.UpdateOne()
				t := models.Target{}
				t.SetID(target.GetID())
				t.FetchByID()
				Expect(target.Title).To(Equal(t.Title))
			})

			It("should soft delete a target record", func() {
				target.SoftDelete()
				t := models.Target{}
				t.SetID(target.GetID())
				t.FetchByID()
				Expect(t.Title).To(Equal(""))
				Expect(t.TargetGroupID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship tests", func() {
		var (
			targetVersion models.TargetVersion
			quiz          models.Quiz
		)

		BeforeEach(func() {
			targetVersion = CreateTargetVersion(target)
			quiz = CreateQuiz(target)
		})

		AfterEach(func() {
			tv := models.TargetVersion{}
			tv.Delete()

			q := models.Quiz{}
			q.Delete()
		})
		Context("TargetGroup Based", func() {
			It("should get the target group", func() {
				target.GetTargetGroup()
				Expect(target.TargetGroup.GetID()).To(Equal(targetGroup.GetID()))
			})
		})

		Context("TargetVersion Based", func() {
			It("should get target versions", func() {
				target.GetVersions()
				Expect(len(target.TargetVersions)).To(Equal(1))
				Expect(target.TargetVersions[0].GetID()).To(Equal(targetVersion.GetID()))
			})
		})

		Context("Quiz Based", func() {
			It("should get quizzes", func() {
				target.GetQuizzes()
				Expect(len(target.Quizzes)).To(Equal(1))
				Expect(target.Quizzes[0].GetID()).To(Equal(quiz.GetID()))
			})
		})
	})
})
