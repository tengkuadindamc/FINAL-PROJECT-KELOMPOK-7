package http_handlers

import (
	"fp-2/dto"
	"fp-2/pkg/errs"
	"fp-2/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) RegisterUser(ctx *gin.Context) {
	var requestBody dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := requestBody.ValidateStruct(); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	createdUser, messageErr := u.userService.RegisterUser(&requestBody)
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (u *userHandler) LoginUser(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	messageErr := requestBody.ValidateStruct()
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	tokenResponse, messageErr := u.userService.LoginUser(&requestBody)

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse)
}

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var requestBody dto.UpdateUserRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	messageErr := requestBody.ValidateStruct()
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	updatedUserResponse, messageErr := u.userService.UpdateUser(userId, &requestBody)

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, updatedUserResponse)

}

func (u *userHandler) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	response, messageErr := u.userService.DeleteUser(userId)

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
