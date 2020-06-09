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
}

// TableName gorm standard table name
func (u *User) TableName() string {
	return userTableName
}

// GetID returns Id of the user
func (u *User) GetID() string {
	return u.ID
}

// SetID sets Id of the user
func (u *User) SetID(id string) {
	u.ID = id
}

// SetCreatedAt sets field createdAt, should only be used in mongodb
func (u *User) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

// SetUpdatedAt sets field UpdatedAt
func (u *User) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

// SetArchivedAt sets field DeletedAt
func (u *User) SetArchivedAt(t *time.Time) {
	u.ArchivedAt = t
}

// UserList defines array of user objects
type UserList []*User

// TableName gorm standard table name
func (u *UserList) TableName() string {
	return userTableName
}

/**
CRUD functions
*/

// Create creates a new user record
func (u *User) Create() error {
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
	err := handler.First(u).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchAll fetchs all Users
func (u *User) FetchAll(ul *UserList) error {
	err := handler.Find(ul).Error
	return err
}

// UpdateOne updates a given user
func (u *User) UpdateOne() error {
	err := handler.Save(u).Error
	return err
}

// Delete deletes user by id
func (u *User) Delete() error {
	err := handler.Delete(u).Error
	return err
}
