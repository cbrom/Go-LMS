package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	courseAuthorTableName = "course_authors"
)

// CourseAuthor defines a model for course authors
type CourseAuthor struct {
	utils.Base
	UserID       string          `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	CourseID     string          `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Course       Course          `gorm:"foreignkey:CourseID"`
	User         User            `gorm:"foreignkey:UserID"`
	Certificates CertificateList `gorm:"foreignkey:Issuer"`
}

// TableName gorm standard table name
func (c *CourseAuthor) TableName() string {
	return courseAuthorTableName
}

// CourseAuthorList defines array of course author objects
type CourseAuthorList []*CourseAuthor

// TableName gorm standard table name
func (c *CourseAuthorList) TableName() string {
	return courseAuthorTableName
}

/**
Relationship functions
*/

// GetCourse returns the Course of this relationship
func (c *CourseAuthor) GetCourse() error {
	return handler.Model(c).Related(&c.Course).Error
}

// GetUser returns the Course Author of this relationship
func (c *CourseAuthor) GetUser() error {
	return handler.Model(c).Related(&c.User).Error
}

/**
CRUD functions
*/

// Create creates a new course author record
func (c *CourseAuthor) Create() error {
	possible := handler.NewRecord(c)
	if possible {
		if err := handler.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches CourseAuthor by id
func (c *CourseAuthor) FetchByID() error {
	err := handler.First(c).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all CourseAuthors
func (c *CourseAuthor) FetchAll(cl *CourseAuthorList) error {
	err := handler.Find(cl).Error
	return err
}

// UpdateOne updates a given course author
func (c *CourseAuthor) UpdateOne() error {
	err := handler.Save(c).Error
	return err
}

// Delete deletes course author by id
func (c *CourseAuthor) Delete() error {
	err := handler.Unscoped().Delete(c).Error
	return err
}

// SoftDelete set's record deleted at field
func (c *CourseAuthor) SoftDelete() error {
	err := handler.Delete(c).Error
	return err
}
