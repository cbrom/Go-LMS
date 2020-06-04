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
	ErrUnableToFetchUserList = errors.New("Unable to fetch user list")
	ErrUnableToUpdateUser    = errors.New("Unable to update user")
	ErrUnableToDeleteUser    = errors.New("Unable to delete user")

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

// FetchAll fetchs all Users
func FetchAll(l *List, query *User) error {
	err := database.Handler().List(l, query)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{"Unknown": []string{err.Error()}}, http.StatusNotFound, ErrUnableToFetchUserList.Error())
		return errs
	}
	return nil
}

// UpdateInfo updates User by given id
func UpdateInfo(u *User, query *UserInfoUpdateRequest) error {

	// check if user exists
	user := &User{
		Name:      query.Name,
		Phone:     query.Phone,
		Title:     query.Title,
		KeySkills: query.KeySkills,
		About:     query.About,
		TimeZone:  query.TimeZone,
	}
	user.SetID(query.ID)
	if err := FetchByID(user); err != nil {
		(err.(*utils.Error)).Message = ErrUnableToFetchUser.Error()
		return err
	}
	err := database.Handler().Update(u, user)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"unknown": []string{err.Error()}}, http.StatusBadRequest, ErrUnableToUpdateUser.Error())
		return errs
	}
	return nil
}

// Delete deletes user by id
func Delete(query *User) error {
	// check if user exists
	if err := FetchByID(query); err != nil {
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
