package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stergiosbamp/go-api/src/auth"
	"github.com/stergiosbamp/go-api/src/dao"
	"github.com/stergiosbamp/go-api/src/models"
	"golang.org/x/crypto/bcrypt"
)

var userDAO = dao.NewUserDAO()
var tokenDAO = dao.NewTokenDAO()

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
	Message string `json:"message"`
	Token   string `json:"token"`
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
	var token auth.AuthProvider

	var userLoginRequest UserLoginRequest

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDAO.FindByUsername(userLoginRequest.Username)
	if err != nil {
		response := UserLoginResponse{
			Message: "User does not exist",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		response := UserLoginResponse{
			Message: "Invalid credentials",
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	tokenString, err := token.GenerateToken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save token of user in database
	tokenDAO.Create(&models.Token{
		UserID: user.ID,
		Token:  tokenString,
		Status: "active",
	})

	response := UserLoginResponse{
		Message: "Login successful",
		Token:   tokenString,
	}

	ctx.JSON(http.StatusOK, response)

}

func Logout(ctx *gin.Context) {
	var token auth.AuthProvider

	tokenString, err := token.ExtractToken(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user (username) from token
	username, err := token.GetUserFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDAO.FindByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save token of user in database
	tokenDAO.UpdateStatus(user, "inactive")

	response := UserLoginResponse{
		Message: "Logout successful",
		Token:   tokenString,
	}

	ctx.JSON(http.StatusOK, response)
}
