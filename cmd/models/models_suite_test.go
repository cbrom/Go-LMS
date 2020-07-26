package models_test

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/migrations"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("Configuration", func() {
	AfterSuite(func() {
		models.CloseDB()
	})

	BeforeSuite(func() {
		db := models.ConnectToTestDatabase()
		migrations.Migrate(db)
	})
})
