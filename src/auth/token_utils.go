package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// ExtractToken extracts the token from the Authorization header.
func ExtractToken(request *http.Request) (string, error) {
	authString := request.Header.Get("Authorization")

	authParts := strings.Split(authString, " ")
	if len(authParts) < 2 {
		return "", errors.New("missing bearer token")
	}

	token := authParts[1]

	return token, nil
}

// GetUsernameFromToken extracts the username from the token.
func GetUsernameFromToken(tokenString string) (string, error) {
	secretKey := conf.GetSecretKey()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["user"].(string), nil
	} else {
		return "", errors.New("invalid token")
	}
}
