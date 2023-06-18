package handler

import (
	"finalproject4/helper"
	"finalproject4/model"
	"finalproject4/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UpdateBalance(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
	authService helper.Service
}

func NewUserHandler(userService service.UserService, authService helper.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (uh *userHandler) RegisterUser(ctx *gin.Context) {
	var userRequest model.UserRegisterRequest
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := uh.userService.RegisterUser(userRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (uh *userHandler) LoginUser(ctx *gin.Context) {
	var userLogin model.UserLoginRequest
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.Login(userLogin)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comparePass := helper.ComparePass([]byte(user.Password), []byte(userLogin.Password))
	if !comparePass {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password is wrong"})
		return
	}

	token, err := uh.authService.GenerateToken(int(user.ID))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": token})
}

func (uh *userHandler) UpdateBalance(ctx *gin.Context) {
	var userBalance model.UserBalanceRequest
	err := ctx.ShouldBindJSON(&userBalance)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet("currentUser").(model.User)
	userID := currentUser.ID

	err = uh.userService.TopUpBalance(userBalance, int(userID))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Balance updated"})
}
