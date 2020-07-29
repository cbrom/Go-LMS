package models_test

import (
	"encoding/json"
	"go-lms-of-pupilfirst/cmd/models"

	"github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContentBlock.Model", func() {
	var (
		course        models.Course
		level         models.Level
		targetGroup   models.TargetGroup
		target        models.Target
		targetVersion models.TargetVersion
		contentBlock  models.ContentBlock
	)

	BeforeEach(func() {
		// init variables
		course = CreateCourse()
		level = CreateLevel(course)
		targetGroup = CreateTargetGroup(level)
		target = CreateTarget(targetGroup)
		targetVersion = CreateTargetVersion(target)
		contentBlock = CreateContentBlock(targetVersion)

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

		cb := models.ContentBlock{}
		cb.Delete()
	})

	Describe("Basic CRUD tests", func() {
		Context("Crud Basics", func() {
			It("should create a content block", func() {
				createdCB := models.ContentBlock{}
				createdCB.SetID(contentBlock.GetID())
				createdCB.FetchByID()
				rawMessage := contentBlock.Content.RawMessage
				var unmarshaledContent struct {
					Key   string
					Value string
				}
				json.Unmarshal(rawMessage, &unmarshaledContent)

				rawMessage = createdCB.Content.RawMessage
				var createdContent struct {
					Key   string
					Value string
				}

				json.Unmarshal(rawMessage, &createdContent)
				Expect(createdContent).To(Equal(unmarshaledContent))
			})

			It("should update an existing content block", func() {
				value := struct {
					Key   string
					Value string
				}{Key: "autorefid", Value: "200"}

				returned, _ := json.Marshal(value)

				contentBlock.Content = postgres.Jsonb{returned}
				contentBlock.UpdateOne()
				cb := models.ContentBlock{}
				cb.SetID(contentBlock.GetID())
				cb.FetchByID()

				rawMessage := contentBlock.Content.RawMessage
				var unmarshaledContent struct {
					Key   string
					Value string
				}
				json.Unmarshal(rawMessage, &unmarshaledContent)

				Expect(value).To(Equal(unmarshaledContent))
			})

			It("should soft delete a target content block", func() {
				contentBlock.SoftDelete()
				cb := models.ContentBlock{}
				cb.SetID(contentBlock.GetID())
				cb.FetchByID()
				Expect(cb.Content).To(Equal(postgres.Jsonb{}))
				Expect(cb.TargetVersionID).To(Equal(""))
			})
		})
	})
})
