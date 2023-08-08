package dao

import (
	"github.com/stergiosbamp/go-api/models"
)

type TokenDAO struct {
	
}

func NewTokenDAO() *TokenDAO {
	return &TokenDAO{}
}

func (t *TokenDAO) Create(token *models.Token) (*models.Token, error) {
	res := db.Create(token)
	return token, res.Error
}

func (t *TokenDAO) UpdateStatus(user *models.User, status string) (*models.Token, error) {
	var token models.Token
	res := db.Model(&token).Where("user_id = ?", user.ID).Update("status", status)
	return &token, res.Error
}

func (t *TokenDAO) FindByTokenString(tokenString string) (*models.Token, error) {
	var token models.Token
	res := db.Where("token = ?", tokenString).First(&token)
	return &token, res.Error
}