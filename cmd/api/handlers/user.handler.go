package handlers

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/auth"
	"log"
	"net/http"
	"strconv"
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
type UserController struct {
}

// SignUp registers user
func SignUp(ctx *gin.Context) {
	// get values
	// build into struct

	var signupBody UserCreateRequest
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

func SignIn(ctx *gin.Context) {
	var signinBody UserLoginRequest
	ctx.BindJSON(&signinBody)
	founduser, err := signinBody.Tologin()
	if err != nil {
		log.Printf("error in user  => %+v", err.Error())
	}
	value := founduser
	token, _ := authenticator.GenerateToken(auth.Claims{})
	c, _ := authenticator.ParseClaims(token)
	// fmt.Println(c)

	ctx.JSON(200, gin.H{
		"message": value,
		"token":   c,
	})
}

// UserLoginRequest spec for login request
type GetusersRequest struct {
	Id int `json:"id"`
}
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}
type UserLoginResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Role           int    `json:"role"`
	Gender         int    `json:"gender"`
	DisabilityType int    `json:"type_of_disability"`
}

// UserCreateRequest spec for signup request
type UserCreateRequest struct {
	Name            string     `json:"name" validate:"required" example:"Groot"`
	Email           string     `json:"email" validate:"required,email,unique" example:"groot@golms.com"`
	Password        string     `json:"password" validate:"required" example:"GrootSecret"`
	PasswordConfirm string     `json:"password_confirm" validate:"required,eqfield=password" example:"GrootSecret"`
	TimeZone        *time.Time `json:"timezone" validate:"required" example:"America/Anchorage"`
}

// ToUser converts UserLoginRequest to User object
func (userLoginRequest *UserLoginRequest) Tologin() (*models.User, error) {

	foundUser := models.User{
		Email: userLoginRequest.Email,
	}
	// UserLogin(ctx context.Context, user model.userluserLoginRequest) (model.UserLoginResponse, error)
	// response, err := controller.UserService.UserLogin(ctx, user)
	//
	foundUser.FetchByEmail()
	if foundUser.GetID() == "" {
		return nil, nil
	}
	saltedPassword := userLoginRequest.Password + foundUser.PasswordSalt
	if err := bcrypt.CompareHashAndPassword(foundUser.PasswordHash, []byte(saltedPassword)); err != nil {
		err = errors.WithStack(errors.New("ErrAuthenticationFailure"))
		return nil, err
	}
	return &foundUser, nil
	//TODO or return the loginresponse

	// if userLoginRequest == nil {
	// 	return nil, errors.New("Null User Request")
	// }
	// user := &models.User{
	// 	Email:    UserLoginRequest.Email,
	// 	password: UserLoginRequest.Password,
	// }
	// return user, nil
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

// TO DO get users
var Users []models.User

func Getuser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := getuserbyid(id)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func Getusers(ctx *gin.Context) {
	ctx.JSON(200, models.UserList{})
}

func getuserbyid(id string) (*models.User, error) {
	for i, u := range Users {
		if u.ID == id {
			return &Users[i], nil
		}
	}
	return nil, errors.New("user not found")
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
