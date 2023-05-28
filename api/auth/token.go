package auth

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/stergiosbamp/go-api/config"
)

const TOKEN_EXP_MINS = 15

type TokenProvider struct {
	conf config.Config
	Token string
}

func (tokenProvider *TokenProvider) GenerateToken(username string) (string, error) {
	secretKey := tokenProvider.conf.GetSecretKey()

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

	tokenProvider.Token = tokenString

	return tokenProvider.Token, nil
}
