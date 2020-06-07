package models

import (
	"go-lms-of-pupilfirst/pkg/database"
	"go-lms-of-pupilfirst/pkg/utils"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	// Database table for User
	userTableName = "users"
)

var (
	// ErrAuthenticationFailure auth failure
	ErrAuthenticationFailure = errors.New("Authentication failed")
	ErrorNotFound            = errors.New("Entity not found")
	ErrForbidden             = errors.New("Attempted action is not allowed")
	ErrUnableToCreateUser    = errors.New("Unable to create User")
	ErrUnableToFetchUser     = errors.New("Unable to fetch user")
	ErrUnableToFetchUserList = errors.New("Unable to fetch user list")
	ErrUnableToUpdateUser    = errors.New("Unable to update user")
	ErrUnableToDeleteUser    = errors.New("Unable to delete user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	ErrResetExpired = errors.New("Reset expired")
)

type User struct {
	utils.Base
	Email                  string     `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,unique,email"`
	PasswordSalt           string     `json:"-"`
	PasswordHash           []byte     `json:"-"`
	Role                   int        `json:"role"`
	SiginInCount           int        `json:"sign_in_count" validate:"omitempty"`
	CurrentSignInAt        *time.Time `json:"current_sign_in_at" validate:"omitempty"`
	LastSignInAt           *time.Time `json:"last_sign_in_at" validate:"omitempty"`
	CurrentSignInIP        string     `json:"-" validate:"omitempty,ip"`
	LastSignInIP           string     `json:"-" validate:"omitempty,ip"`
	RememberToken          string     `json:"remember_token" validate:"omitempty"`
	ConfirmedAt            *time.Time `json:"confirmed_at" validate:"omitempty"`
	ConfirmationMailSentAt *time.Time `json:"confirmation_mail_sent_at" validate:"omitempty"`
	Name                   string     `json:"name" validate:"min=3,max=10,omitempty"`
	Phone                  string     `json:"phone" validate:"omitempty"`
	Title                  string     `json:"title" validate:"omitempty"`
	KeySkills              string     `json:"key_skills" validate:"omitempty"`
	About                  string     `gorm:"type:text" json:"about" validate:"omitempty"`

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

// List defines array of user objects
type List []*User

// TableName gorm standard table name
func (u *List) TableName() string {
	return userTableName
}

/**
CRUD functions
*/

func (u *User) Create() error {
	err := database.Handler().Insert(u)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"unknown": []string{err.Error()}}, http.StatusBadRequest, ErrUnableToCreateUser.Error())
		return errs
	}
	return nil
}

// FetchByID fetches User by id
func (u *User) FetchByID() error {
	err := database.Handler().One(u)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"user_id": []string{err.Error()}}, http.StatusNotFound, ErrUnableToFetchUser.Error())
		return errs
	}
	return nil
}

// FetchAll fetchs all Users
func (query *User) FetchAll(l *List) error {
	err := database.Handler().List(l, query)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{"Unknown": []string{err.Error()}}, http.StatusNotFound, ErrUnableToFetchUserList.Error())
		return errs
	}
	return nil
}

// Delete deletes user by id
func (query *User) Delete() error {
	// check if user exists
	if err := query.FetchByID(); err != nil {
		(err.(*utils.Error)).Message = ErrUnableToFetchUser.Error()
		return err
	}
	err := database.Handler().Remove(query)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"unknown": []string{err.Error()}}, http.StatusBadRequest, ErrUnableToDeleteUser.Error())
		return errs
	}

	return nil
}
