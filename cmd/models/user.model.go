package models

import (
	"go-lms-of-pupilfirst/pkg/utils"
	"time"
)

const (
	// Database table for User
	userTableName = "users"
)

// User struct for users table
type User struct {
	utils.Base
	Email                  string `gorm:"type:varchar(100);unique_index" `
	PasswordSalt           string
	PasswordHash           []byte
	Role                   int
	SiginInCount           int
	CurrentSignInAt        *time.Time
	LastSignInAt           *time.Time
	CurrentSignInIP        string
	LastSignInIP           string
	RememberToken          string
	ConfirmedAt            *time.Time
	ConfirmationMailSentAt *time.Time
	Name                   string
	Phone                  string
	Title                  string
	KeySkills              string
	About                  string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`

	AuthoredCourses CourseAuthorList  `gorm:"foreignkey:UserID"`
	Courses         StudentCourseList `gorm:"foreignkey:UserID"`
}

// TableName gorm standard table name
func (u *User) TableName() string {
	return userTableName
}

// UserList defines array of user objects
type UserList []*User

// TableName gorm standard table name
func (u *UserList) TableName() string {
	return userTableName
}

/**
Relationship functions
*/

// GetAuthoredCourses returns a list of authored courses
func (u *User) GetAuthoredCourses() error {
	return handler.Model(u).Related(&u.AuthoredCourses).Error
}

// GetCourses returns a list of courses from a list of authored courses
func (u *User) GetCourses() error {
	return handler.Model(u).Related(&u.Courses).Error
}

/**
CRUD functions
*/

// Create creates a new user record
func (u *User) Create() error {
	if handler == nil {
		return errHandlerNotSet
	}
	possible := handler.NewRecord(u)
	if possible {
		if err := handler.Create(u).Error; err != nil {
			return err
		}
	}

	return nil
}

// FetchByID fetches User by id
func (u *User) FetchByID() error {
	if handler == nil {
		return errHandlerNotSet
	}
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchByEmail fetches User by email
func (u *User) FetchByEmail() error {
	if handler == nil {
		return errHandlerNotSet
	}
	err := handler.Where("email = ?", u.Email).First(&u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Users
func (u *User) FetchAll(ul *UserList) error {
	if handler == nil {
		return errHandlerNotSet
	}
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given user
func (u *User) UpdateOne() error {
	if handler == nil {
		return errHandlerNotSet
	}
	err := handler.Save(u).Error
	return err
}

// Delete deletes user by id
func (u *User) Delete() error {
	if handler == nil {
		return errHandlerNotSet
	}
	err := handler.Unscoped().Delete(u).Error
	return err
}

// SoftDelete set's deleted at date
func (u *User) SoftDelete() error {
	if handler == nil {
		return errHandlerNotSet
	}

	err := handler.Delete(u).Error
	return err
}
