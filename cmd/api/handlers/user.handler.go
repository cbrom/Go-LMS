package handlers

import (
	"go-lms-of-pupilfirst/cmd/models/user"
	"go-lms-of-pupilfirst/pkg/auth"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) SignUp(ctx *gin.Context) {
	// get values
	// build into struct

	var signupBody user.UserCreateRequest
	ctx.BindJSON(&signupBody)
	usr, err := signupBody.GetUser()
	if err != nil {
		log.Printf("error in user get => %+v", err.Error())
	}
	value := user.Create(usr)
	token, _ := authenticator.GenerateToken(auth.Claims{})

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   token,
	})
}
