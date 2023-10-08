package dao

import (
	"github.com/stergiosbamp/go-api/models"
)

type UserDAO struct {
	
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (u *UserDAO) Create(user *models.User) (*models.User, error) {
	res := db.Create(user)
	return user, res.Error
}

func (u *UserDAO) FindByUsername(username string) (*models.User, error) {
	var user models.User
	res := db.Where("username = ?", username).First(&user)
	return &user, res.Error
}