package middleware

import (
	"bookhub/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cleintAccessToken := ctx.Request.Header.Get("Authentication")
		if cleintAccessToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(cleintAccessToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.FirstName)
		ctx.Set("last_name", claims.LastName)
		ctx.Set("uid", claims.Uid)
		ctx.Set("user_type", claims.UserType)
		ctx.Next()
	}
}
