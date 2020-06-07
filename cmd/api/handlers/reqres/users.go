package reqres

import (
	"net/http"
	"time"

	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/database"
	"go-lms-of-pupilfirst/pkg/utils"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// UserLoginRequest spec for login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// UserCreateRequest spec for signup request
type UserCreateRequest struct {
	Name            string     `json:"name" validate:"required" example:"Groot"`
	Email           string     `json:"email" validate:"required,email,unique" example:"groot@golms.com"`
	Password        string     `json:"password" validate:"required" example:"GrootSecret"`
	PasswordConfirm string     `json:"password_confirm" validate:"required,eqfield=password" example:"GrootSecret"`
	TimeZone        *time.Time `json:"timezone" validate:"required" example:"America/Anchorage"`
}

// ToUser converts UserCreateRequest to User object
func (userCreateRequest *UserCreateRequest) ToUser() (*models.User, error) {
	if userCreateRequest == nil {
		return nil, errors.New("Null User Create Request")
	}

	passwordSalt := uuid.NewRandom().String()
	saltedPassword := userCreateRequest.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Error generating password hash")
	}

	user := &models.User{
		Name:         userCreateRequest.Name,
		Email:        userCreateRequest.Email,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		TimeZone:     userCreateRequest.TimeZone,
	}
	return user, nil
}

// UserInfoUpdateRequest - spec for updating user info
type UserInfoUpdateRequest struct {
	ID        string `json:"id" validate:"required,uuid" example:"c01bdef7-173f-4d29-3edc60baf6a2"`
	Name      string `json:"name" validate:"min=3,max=10,omitempty"`
	Phone     string `json:"phone" validate:"omitempty"`
	Title     string `json:"title" validate:"omitempty"`
	KeySkills string `json:"key_skills" validate:"omitempty"`
	About     string `gorm:"type:text" json:"about" validate:"omitempty"`

	TimeZone *time.Time `json:"timezone" validation:"omitempty"`
}

// UpdateInfo updates User by given id
func UpdateInfo(u *models.User, query *UserInfoUpdateRequest) error {

	// check if user exists
	user := &models.User{
		Name:      query.Name,
		Phone:     query.Phone,
		Title:     query.Title,
		KeySkills: query.KeySkills,
		About:     query.About,
		TimeZone:  query.TimeZone,
	}
	user.SetID(query.ID)
	if err := user.FetchByID(); err != nil {
		(err.(*utils.Error)).Message = models.ErrUnableToFetchUser.Error()
		return err
	}
	err := database.Handler().Update(u, user)
	if err != nil {
		errs := utils.GenerateError(map[string][]string{
			"unknown": []string{err.Error()}}, http.StatusBadRequest, models.ErrUnableToUpdateUser.Error())
		return errs
	}
	return nil
}
