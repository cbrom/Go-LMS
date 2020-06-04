package handlers

import (
	"go-lms-of-pupilfirst/pkg/auth"

	"github.com/gin-gonic/gin"
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
func ApplyRoutes(r *gin.Engine, authenticator *auth.Authenticator) {
	authenticator = authenticator
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/ping", pingHandler)
	}
}
