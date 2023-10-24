package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stergiosbamp/go-api/src/config"
	"github.com/stergiosbamp/go-api/src/dao"
)

const TOKEN_EXP_MINS = 15

type AuthProvider struct {
	conf  config.Config
	Token string
	dao.TokenDAO
}

func (authProvider *AuthProvider) GenerateToken(username string) (string, error) {
	secretKey := authProvider.conf.GetSecretKey()

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

	authProvider.Token = tokenString

	return authProvider.Token, nil
}

func (authProvider *AuthProvider) Authenticate() (bool, error) {
	secretKey := authProvider.conf.GetSecretKey()

	tokenString := authProvider.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

func (authProvider *AuthProvider) ExtractToken(request *http.Request) (string, error) {
	authString := request.Header.Get("Authorization")

	authParts := strings.Split(authString, " ")
	if len(authParts) < 2 {
		return "", errors.New("missing bearer token")
	}

	token := authParts[1]

	return token, nil
}

func (authProvider *AuthProvider) GetUserFromToken(tokenString string) (string, error) {
	secretKey := authProvider.conf.GetSecretKey()

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

func (authProvider *AuthProvider) IsTokenBlacklisted() bool {
	tokenString := authProvider.Token
	token, err := authProvider.TokenDAO.FindByTokenString(tokenString)

	if err != nil {
		return true
	}

	if token.Status == "inactive" {
		return true
	}

	return false
}
