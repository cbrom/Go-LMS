package routes

import (
	"go-lms-of-pupilfirst/cmd/api/handlers"
	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticator *auth.Authenticator
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, auth *auth.Authenticator, DB *gorm.DB, UserController *handlers.UserController) {
	models.SetRepoDB(DB)
	authenticator = auth
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/ping", pingHandler)
		apiV1.POST("/register", UserController.SignUp)
		apiV1.GET("/verifyemail/:verificationCode", UserController.VerifyEmail)

	}
}
