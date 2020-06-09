package handlers

import (
	"go-lms-of-pupilfirst/cmd/api/handlers/reqres"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/auth"
	"log"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	errAuthenticationFailure = errors.New("Authentication failed")
	errorNotFound            = errors.New("Entity not found")
	errForbidden             = errors.New("Attempted action is not allowed")
	errUnableToCreateUser    = errors.New("Unable to create User")
	errUnableToFetchUser     = errors.New("Unable to fetch user")
	errUnableToFetchUserList = errors.New("Unable to fetch user list")
	errUnableToUpdateUser    = errors.New("Unable to update user")
	errUnableToDeleteUser    = errors.New("Unable to delete user")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	ErrResetExpired = errors.New("Reset expired")
)

// UserController is an anonymous struct for user controller
type UserController struct{}

// SignUp registers user
func (ctrl *UserController) SignUp(ctx *gin.Context) {
	// get values
	// build into struct

	var signupBody reqres.UserCreateRequest
	ctx.BindJSON(&signupBody)
	usr, err := signupBody.ToUser()
	if err != nil {
		log.Printf("error in user get => %+v", err.Error())
	}
	value := usr.Create()
	token, _ := authenticator.GenerateToken(auth.Claims{})

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   token,
	})
}

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
