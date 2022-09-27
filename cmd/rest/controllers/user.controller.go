package controllers

import (
	"fmt"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/auth"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	authenticator *auth.Authenticator
)

// UserCreateRequest spec for signup request
type UserCreateRequest struct {
	Name            string     `json:"name" validate:"required" example:"Groot"`
	Email           string     `json:"email" validate:"required,email,unique" example:"groot@golms.com"`
	Password        string     `json:"password" validate:"required" example:"GrootSecret"`
	PasswordConfirm string     `json:"password_confirm" validate:"required,eqfield=password" example:"GrootSecret"`
	TimeZone        *time.Time `json:"timezone" validate:"required" example:"America/Anchorage"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

// GetTimeFromStamp changes timestamp string to  *time.Time
func GetTimeFromStamp(ts string) *time.Time {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil
	}
	tm := time.Unix(i, 0)
	return &tm
}

func Signup(c *gin.Context) {
	var signupBody UserCreateRequest

	err := c.BindJSON(&signupBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json",
		})
		c.Abort()
		return
	}

	// HashPassword encrypts user password
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := signupBody.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error generating password hash",
		})
	}

	usr := &models.User{
		Name:         signupBody.Name,
		Email:        signupBody.Email,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		TimeZone:     signupBody.TimeZone,
	}

	err = usr.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
	}

	c.JSON(http.StatusOK, usr)
}

// UserLoginRequest spec for login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {

	var loginBody UserLoginRequest

	err := c.Bind(&loginBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json",
		})
		c.Abort()
		return
	}

	foundUser := &models.User{
		Email: loginBody.Email,
	}

	foundUser.FetchByEmail()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Failed to get user",
		})
		c.Abort()
		return
	}

	saltedPassword := loginBody.Password + foundUser.PasswordSalt
	if err := bcrypt.CompareHashAndPassword(foundUser.PasswordHash, []byte(saltedPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid user credentials",
		})
		c.Abort()
		return
	}

	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   foundUser.ID,
			Audience:  "",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 20000).Unix(),
		},
	}

	token, err := authenticator.GenerateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error signing token",
		})
		c.Abort()
		return
	}

	cl, _ := authenticator.ParseClaims(token)
	fmt.Println(cl)

	// // set Token in cookies
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", token, 20000 "", "", false, true)

	tokenResponse := LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, tokenResponse)

}

// Return the Users
func GetUser(c *gin.Context) {
	idQuery := c.Param("id")

	usr := &models.User{}
	usr.SetID(idQuery)
	err := usr.FetchByID()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "User ID not provided",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": usr,
	})
}
