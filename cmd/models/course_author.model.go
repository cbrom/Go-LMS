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
	UserID   string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	CourseID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
}

// TableName gorm standard table name
func (u *CourseAuthor) TableName() string {
	return courseAuthorTableName
}

// CourseAuthorList defines array of course author objects
type CourseAuthorList []*CourseAuthor

// TableName gorm standard table name
func (u *CourseAuthorList) TableName() string {
	return courseAuthorTableName
}

/**
CRUD functions
*/

// Create creates a new course author record
func (u *CourseAuthor) Create() error {
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches CourseAuthor by id
func (u *CourseAuthor) FetchByID() error {
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all CourseAuthors
func (u *CourseAuthor) FetchAll(ul *CourseAuthorList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given course author
func (u *CourseAuthor) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes course author by id
func (u *CourseAuthor) Delete() error {
	err := handler.Delete(u).Error
	return err
}
