package handlers

import (
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/pkg/middlewares"
	"go-lms-of-pupilfirst/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

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

// CreateUser godoc
// @Summary Registers a user
// @Description creates user directory
// @Tags Users
// @Param Body body UserCreateRequest true "The body to create a user"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /register [post]
// [...]Registor User
func (ctrl *UserController) SignUp(ctx *gin.Context) {
	// get values
	// build into struct

	var signupBody UserCreateRequest
	if err := ctx.ShouldBindJSON(&signupBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if signupBody.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Email field is required"})
		return
	}
	if signupBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Password field is required"})
		return
	}
	if signupBody.Password != signupBody.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}
	passwordSalt := uuid.NewRandom().String()
	saltedPassword := signupBody.Password + passwordSalt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user := models.User{
		Name:         signupBody.Name,
		Email:        strings.ToLower(signupBody.Email),
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		Verified:     false,
		// TimeZone:     signupBody.TimeZone,
	}

	result := user.Create()

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": result})

	config, _ := configs.LoadConfig()
	// Generate Verification Code
	code := randstr.String(20)

	verification_code := utils.Encode(code)

	// Update User in Database
	user.VerificationCode = verification_code
	user.Save()
	var firstName = user.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// ? Send Email
	emailData := middlewares.EmailData{
		URL:       config.ClientOrigin + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	middlewares.SendEmail(&user, &emailData)
	message := "We sent an email with a verification code to " + user.Email
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})

}

// VerifyEmail godoc
// @Summary Verify a user
// @Description verificationCode varifay a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param verification_code path string true "verificationCode"
// @Success 200  {object} string "success"
// @Failure 400  {object} string "error"
// @Router /verifyemail/{verification_code} [get]
// [...] Verify Email
func (ctrl *UserController) VerifyEmail(ctx *gin.Context) {

	code := ctx.Param("verificationCode")
	verification_code := utils.Encode(code)
	updatedUser := &models.User{}
	updatedUser.VerificationCode = verification_code
	if updatedUser.FetchByVerificationCode() != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
		return
	}

	if updatedUser.Verified {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified"})
		return
	}

	updatedUser.VerificationCode = ""
	updatedUser.Verified = true
	updatedUser.Save()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email verified successfully"})
}

// UserLoginRequest spec for login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
}

// UserCreateRequest spec for signup request
type UserCreateRequest struct {
	Name            string `json:"name" validate:"required" example:"Groot"`
	Email           string `json:"email" validate:"required,email,unique" example:"groot@golms.com"`
	Password        string `json:"password" validate:"required" example:"GrootSecret"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=password" example:"GrootSecret"`
	// TimeZone        *time.Time `json:"timezone" validate:"required" example:"America/Anchorage"`
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
