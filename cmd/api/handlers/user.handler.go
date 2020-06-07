package handlers

import (
	"go-lms-of-pupilfirst/cmd/api/handlers/reqres"
	"go-lms-of-pupilfirst/pkg/auth"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

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
