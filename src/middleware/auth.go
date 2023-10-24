package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/auth"
)

var authProvider = auth.AuthProvider{}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := authProvider.ExtractToken(ctx.Request)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": err.Error()})
			ctx.Abort()
			return
		}

		// set token
		authProvider.Token = token
		verified, err := authProvider.Authenticate()

		if !verified {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": err.Error()})
			ctx.Abort()
			return
		}

		// is token blacklisted (after logout)
		blacklisted := authProvider.IsTokenBlacklisted()
		if blacklisted {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": "token is blacklisted"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
