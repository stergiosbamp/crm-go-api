package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/auth"
)

func JwtAuthFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token provided?
		token, err := auth.ExtractToken(ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": "Missing token"})
			ctx.Abort()
			return
		}

		// token blacklisted (logged out) ?
		var tokenRevoker auth.TokenRevoker = auth.NewRedisTokenRevoker(ctx)
		if tokenRevoker.IsTokenRevoked(token) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": "Token is revoked"})
			ctx.Abort()
			return
		}

		// token valid ?
		var authProvider = auth.NewAuthProvider()
		authenticated, err := authProvider.Authenticate(token)
		if !authenticated {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "error": "Invalid token"})
			ctx.Abort()
			return
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Internal error", "error": err.Error()})
			ctx.Abort()
			return
		}
		
		ctx.Next()
	}
}
