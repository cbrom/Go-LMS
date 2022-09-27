package routes

import (
	"go-lms-of-pupilfirst/cmd/rest/controllers"
	"go-lms-of-pupilfirst/pkg/auth"
	"go-lms-of-pupilfirst/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	authenticator *auth.Authenticator
)

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, auth *auth.Authenticator) {
	// models.SetRepoDB(db)
	authenticator = auth

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/profile", middlewares.JWTAuthMiddleware(authenticator), controllers.GetUser)

}
