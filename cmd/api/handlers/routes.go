package handlers

import (
	"go-lms-of-pupilfirst/pkg/auth"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to gin engine
func ApplyRoutes(r *gin.Engine, authenticator *auth.Authenticator) {
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/ping", pingHandler)
	}
}
