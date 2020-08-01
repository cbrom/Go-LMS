package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
)

var (
	studentCourseTableName = "student_courses"
)

// StudentCourse defines a model for course students
type StudentCourse struct {
	utils.Base
	UserID   string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	CourseID string `sql:"type:uuid;" validate:"omitempty,uuid,required"`
	Course   Course `gorm:"foreignkey:CourseID"`
	User     User   `gorm:"foreignkey:UserID"`
}

// TableName gorm standard table name
func (s *StudentCourse) TableName() string {
	return studentCourseTableName
}

// StudentCourseList defines array of course student objects
type StudentCourseList []*StudentCourse

// TableName gorm standard table name
func (s *StudentCourseList) TableName() string {
	return studentCourseTableName
}

/**
Relationship functions
*/

// GetCourse returns student course
func (s *StudentCourse) GetCourse() error {
	return handler.Model(s).Related(&s.Course).Error
}

// GetUser returns student course
func (s *StudentCourse) GetUser() error {
	return handler.Model(s).Related(&s.User).Error
}

/**
CRUD functions
*/

// Create creates a new course student record
func (s *StudentCourse) Create() error {
	possible := handler.NewRecord(s)
	if possible {
		if err := handler.Create(s).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches StudentCourse by id
func (s *StudentCourse) FetchByID() error {
	err := handler.First(s).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all StudentCourses
func (s *StudentCourse) FetchAll(sl *StudentCourseList) error {
	err := handler.Find(sl).Error
	return err
}

// UpdateOne updates a given course student
func (s *StudentCourse) UpdateOne() error {
	err := handler.Save(s).Error
	return err
}

// Delete deletes course student by id
func (s *StudentCourse) Delete() error {
	err := handler.Unscoped().Delete(s).Error
	return err
}

// SoftDelete sets deleted at field
func (s *StudentCourse) SoftDelete() error {
	return handler.Delete(s).Error
}
