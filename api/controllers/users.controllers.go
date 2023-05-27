package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/dao"
	"github.com/stergiosbamp/go-api/models"
	"golang.org/x/crypto/bcrypt"
)

var userDAO = dao.NewUserDAO()

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}


func Register(ctx *gin.Context) {
	var userRegisterRequest UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userRegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.Username = userRegisterRequest.Username
	user.Password = string(hashedPassword)

	userCreated, err := userDAO.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userRegisterResponse UserRegisterResponse
	userRegisterResponse.ID = userCreated.ID
	userRegisterResponse.Username = userCreated.Username

	ctx.JSON(http.StatusOK, userRegisterResponse)
}

func Login(ctx *gin.Context) {
	var userLoginRequest UserLoginRequest

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDAO.FindByUsername(userLoginRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}