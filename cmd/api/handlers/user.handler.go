package handlers

import (
	"errors"
	"go-lms-of-pupilfirst/cmd/api/handlers/reqres"
	"go-lms-of-pupilfirst/pkg/auth"
	"log"

	"github.com/gin-gonic/gin"
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
