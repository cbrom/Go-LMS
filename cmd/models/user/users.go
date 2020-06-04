package user

import (
	"go-lms-of-pupilfirst/pkg/database"
	"go-lms-of-pupilfirst/pkg/utils"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrAuthenticationFailure auth failure
	ErrAuthenticationFailure = errors.New("Authentication failed")
	ErrorNotFound            = errors.New("Entity not found")
	ErrForbidden             = errors.New("Attempted action is not allowed")
	ErrUnableToCreateUser    = errors.New("Unable to create User")
	ErrUnableToFetchUser     = errors.New("Unable to fetch user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	ErrResetExpired = errors.New("Reset expired")
)

// GetUserFromCreateRequest returns user object from request
func (user *UserCreateRequest) GetUserFromCreateRequest() (*User, error) {
	if user == nil {
		return nil, nil
	}

	usr := &User{
		Name: user.Name,
	}

	return usr, nil
}

// Create creates a new User record of given values
func Create(u *User) error {
	err := database.Handler().Insert(u)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"unknown": []string{err.Error()}}, http.StatusBadRequest, ErrUnableToCreateUser.Error())
		return errs
	}
	return nil
}

// FetchByID fetches User by id
func FetchByID(u *User) error {
	err := database.Handler().One(u)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"user_id": []string{err.Error()}}, http.StatusNotFound, ErrUnableToFetchUser.Error())
		return errs
	}
	return nil
}
