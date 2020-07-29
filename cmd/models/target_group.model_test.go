package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetGroup.Model", func() {
	var (
		course      models.Course
		level       models.Level
		targetGroup models.TargetGroup
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
	})

	AfterEach(func() {
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		l := models.Level{}
		l.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a new target group", func() {
				createdTG := models.TargetGroup{}
				createdTG.SetID(targetGroup.GetID())
				createdTG.FetchByID()
				Expect(createdTG.Name).To(Equal(targetGroup.Name))
			})

			It("should update an existing target group", func() {
				targetGroup.Name = "Changed Target Group Name"
				targetGroup.UpdateOne()
				t := models.TargetGroup{}
				t.SetID(targetGroup.GetID())
				t.FetchByID()
				Expect(targetGroup.Name).To(Equal(t.Name))
			})

			It("should soft delete a target group record", func() {
				targetGroup.SoftDelete()
				t := models.TargetGroup{}
				t.SetID(targetGroup.GetID())
				t.FetchByID()
				Expect(t.Name).To(Equal(""))
				Expect(t.LevelID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship tests", func() {
		var (
			target models.Target
		)

		BeforeEach(func() {
			target = CreateTarget(targetGroup)
		})

		AfterEach(func() {
			t := models.Target{}
			t.Delete()
		})
		Context("Level Based", func() {
			It("should get the level", func() {
				targetGroup.GetLevel()
				Expect(targetGroup.Level.GetID()).To(Equal(level.GetID()))
			})
		})

		Context("Target Based", func() {
			It("should get targets", func() {
				targetGroup.GetTargets()
				Expect(len(targetGroup.Targets)).To(Equal(1))
				Expect(targetGroup.Targets[0].GetID()).To(Equal(target.GetID()))
			})
		})
	})
})
