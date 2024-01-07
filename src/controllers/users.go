package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/stergiosbamp/go-api/src/auth"
	"github.com/stergiosbamp/go-api/src/dao"
	"github.com/stergiosbamp/go-api/src/models"
)

var userDAO = dao.NewUserDAO()
var authProvider = auth.NewAuthProvider()

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

type UserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
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
		response := UserResponse{
			Message: "User does not exist",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		response := UserResponse{
			Message: "Invalid credentials",
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	tokenString, err := authProvider.GenerateToken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		Message: "Login successful",
		Token:   tokenString,
	}

	ctx.JSON(http.StatusOK, response)

}

func Logout(ctx *gin.Context) {
	tokenString, err := auth.ExtractToken(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Revoke user's session
	var tokenRevoker auth.TokenRevoker = auth.NewRedisTokenRevoker(ctx)
	tokenRevoker.RevokeToken(tokenString)

	response := UserResponse{
		Message: "Logout successful",
	}

	ctx.JSON(http.StatusOK, response)
}
