package handlers

import (
	// "go-lms-of-pupilfirst/Go-LMS/cmd/api/handlers"
	"go-lms-of-pupilfirst/pkg/auth"

	"go-lms-of-pupilfirst/cmd/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticator *auth.Authenticator
)

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, auth *auth.Authenticator, db *gorm.DB) {

	models.SetRepoDB(db)
	authenticator = auth
	apiV1 := r.Group("/v1")
	{
		apiV1.POST("/users/signup", SignUp)    // done
		apiV1.POST("/users/signin", SignIn)    //done
		apiV1.GET("/users/:id", Getuser)       // done
		apiV1.GET("/users/", Getusers)         //done
		apiV1.PUT("/users/:id", updateuser)    //done
		apiV1.DELETE("/users/:id", delateuser) //done

		// apiV1.GET("/users/all", Getusers)

	}
}
