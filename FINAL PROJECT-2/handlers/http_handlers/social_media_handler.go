package http_handlers

import (
	"fp-2/dto"
	"fp-2/pkg/errs"
	"fp-2/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService services.SocialMediaService) *socialMediaHandler {
	return &socialMediaHandler{socialMediaService: socialMediaService}
}

func (sm *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var requestBody dto.CreateSocialMediaRequest

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

	newSMResponse, messageErr := sm.socialMediaService.CreateSocialMedia(&requestBody, userId)
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, newSMResponse)
}

func (sm *socialMediaHandler) GetAllSocialMedias(ctx *gin.Context) {
	allSMResponse, messageErr := sm.socialMediaService.GetAllSocialMedias()

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, allSMResponse)
}

func (sm *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		idError := errs.NewBadRequest(err.Error())
		ctx.JSON(idError.StatusCode(), idError)
		return
	}
	if socialMediaId < 0 {
		idError := errs.NewBadRequest("Social Media ID value must be positive")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.CreateSocialMediaRequest

	err = ctx.ShouldBindJSON(&requestBody)
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

	updatedSMResponse, messageErr := sm.socialMediaService.UpdateSocialMedia(uint(socialMediaId), &requestBody)

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, updatedSMResponse)
}

func (sm *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		idError := errs.NewBadRequest(err.Error())
		ctx.JSON(idError.StatusCode(), idError)
		return
	}
	if socialMediaId < 0 {
		idError := errs.NewBadRequest("Social Media ID value must be positive")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	deletedSMResponse, messageErr := sm.socialMediaService.DeleteSocialMedia(uint(socialMediaId))

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, deletedSMResponse)
}
