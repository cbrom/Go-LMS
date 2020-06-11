package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"go-lms-of-pupilfirst/pkg/auth"
)

// JWTAuthMiddleware gets token from header and sets claims and auth key in the context
func JWTAuthMiddleware(authenticator *auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")

		// failed to read cookie
		if err != nil {
			// try Http Header
			authorization := ctx.Request.Header.Get("Authorization")
			if authorization != "" {
				sp := strings.Split(authorization, "Bearer ")
				// invalid token
				if len(sp) >= 1 {
					tokenString = sp[1]
				}
			}

		}

		claims, _ := authenticator.ParseClaims(tokenString)

		ctx.Set("claims", claims)
		ctx.Set("auth_key", auth.Key)
		ctx.Next()
	}
}
