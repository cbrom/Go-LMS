package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

var (
	courseTableName = "courses"
)

// Course is a model for Courses table
type Course struct {
	utils.Base
	Name                string `gorm:"type:varchar(100);unique" `
	EndsAt              *time.Time
	Description         string
	EnableLeadboard     bool
	PublicSignup        bool
	Featured            bool
	About               string `gorm:"type:varchar(100)"`
	ProgressionBehavior string `gorm:"type:varchar(100)"`
	ProgressionLimit    int
	Certificates        CertificateList        `gorm:"foreignkey:CourceID"`
	Authors             CourseAuthorList       `gorm:"foreignkey:CourseID"`
	EvaluationCriterias EvaluationCriteriaList `gorm:"foreignkey:CourseID"`
	Levels              LevelList              `gorm:"foreignkey:CourseID"`
	Students            StudentCourseList      `gorm:"foreignkey:CourseID"`
}

// TableName gorm standard table name
func (c *Course) TableName() string {
	return courseTableName
}

// CourseList defines array of course objects
type CourseList []*Course

// TableName gorm standard table name
func (c *CourseList) TableName() string {
	return courseTableName
}

/**
CRUD functions
*/

// Create creates a new course record
func (c *Course) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Course by id
func (c *Course) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Courses
func (c *Course) FetchAll(cl *CourseList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given course
func (c *Course) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes course by id
func (c *Course) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's deleted at date
func (c *Course) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
