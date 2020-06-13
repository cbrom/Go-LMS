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
}

// TableName gorm standard table name
func (u *Course) TableName() string {
	return courseTableName
}

// CourseList defines array of course objects
type CourseList []*Course

// TableName gorm standard table name
func (u *CourseList) TableName() string {
	return courseTableName
}

/**
CRUD functions
*/

// Create creates a new course record
func (u *Course) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches Course by id
func (u *Course) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Courses
func (u *Course) FetchAll(ul *CourseList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given course
func (u *Course) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes course by id
func (u *Course) Delete() error {
	err := handler.Delete(u).Error
	return err
}
