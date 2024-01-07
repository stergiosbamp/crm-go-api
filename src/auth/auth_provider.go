package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stergiosbamp/go-api/src/config"
)

const TOKEN_EXP_MINS = 15 // token expiration time in minutes

var conf = config.NewConfig()

type AuthProvider struct{}

func NewAuthProvider() *AuthProvider {
	return &AuthProvider{}
}

func (authProvider *AuthProvider) GenerateToken(username string) (string, error) {
	secretKey := conf.GetSecretKey()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * TOKEN_EXP_MINS).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = username

	// sign token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (authProvider *AuthProvider) Authenticate(tokenStr string) (bool, error) {
	secretKey := conf.GetSecretKey()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return true, nil
	} else {
		return false, errors.New("invalid token")
	}
}

