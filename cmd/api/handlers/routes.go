package handlers

import (
	// "go-lms-of-pupilfirst/Go-LMS/cmd/api/handlers"
	"go-lms-of-pupilfirst/pkg/auth"

	"go-lms-of-pupilfirst/cmd/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// func (ctrl *UserController) setupRouter() {
// 	router := gin.Default()

// 	router.POST("/users", ctrl.SignUp)
// 	router.POST("/users/login", ctrl.SignIn)
// 	// router.POST("/tokens/renew_access", server.renewAccessToken)

// 	ctrl.r = router
// }

var (
	authenticator *auth.Authenticator
)

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, auth *auth.Authenticator, db *gorm.DB) {

	models.SetRepoDB(db)
	authenticator = auth
	apiV1 := r.Group("/v1")
	{
		apiV1.POST("/users", SignIn)       // done
		apiV1.POST("/users/login", SignUp) //done
		apiV1.GET("/users/:id", Getuser)   // done
		apiV1.GET("/users/all", Getusers)
		
	}
}
