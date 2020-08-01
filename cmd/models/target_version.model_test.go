package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetVersion.Model", func() {
	var (
		course        models.Course
		level         models.Level
		targetGroup   models.TargetGroup
		target        models.Target
		targetVersion models.TargetVersion
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)
		targetVersion = CreateTargetVersion(target)

	})

	AfterEach(func() {
		u := models.User{}
		u.Delete()

		c := models.Course{}
		c.Delete()

		l := models.Level{}
		l.Delete()

		tg := models.TargetGroup{}
		tg.Delete()

		t := models.Target{}
		t.Delete()

		tv := models.TargetVersion{}
		tv.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a target version", func() {
				createdTV := models.TargetVersion{}
				createdTV.SetID(targetVersion.GetID())
				createdTV.FetchByID()
				Expect(createdTV.VersionName).To(Equal(targetVersion.VersionName))
			})

			It("should update an existing target version", func() {
				targetVersion.VersionName = "Changed Target VersionName"
				targetVersion.UpdateOne()
				tv := models.TargetVersion{}
				tv.SetID(targetVersion.GetID())
				tv.FetchByID()
				Expect(targetVersion.VersionName).To(Equal(tv.VersionName))
			})

			It("should soft delete a target version record", func() {
				targetVersion.SoftDelete()
				tv := models.TargetVersion{}
				tv.SetID(targetVersion.GetID())
				tv.FetchByID()
				Expect(tv.VersionName).To(Equal(""))
				Expect(tv.TargetID).To(Equal(""))
			})
		})
	})

	Describe("Basic Relationship tests", func() {
		var (
			contentBlock models.ContentBlock
		)

		BeforeEach(func() {
			contentBlock = CreateContentBlock(targetVersion)
		})

		AfterEach(func() {
			cb := models.ContentBlock{}
			cb.Delete()
		})

		Context("Target Based", func() {
			It("should get target", func() {
				targetVersion.GetTarget()
				Expect(targetVersion.Target.GetID()).To(Equal(target.GetID()))
			})
		})

		Context("ContentBlock Based", func() {
			It("should get content blocks", func() {
				targetVersion.GetContentBlocks()
				Expect(len(targetVersion.ContentBlocks)).To(Equal(1))
				Expect(targetVersion.ContentBlocks[0].GetID()).To(Equal(contentBlock.GetID()))
			})
		})
	})
})
