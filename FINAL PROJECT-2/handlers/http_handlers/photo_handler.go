package http_handlers

import (
	"fp-2/dto"
	"fp-2/entities"
	"fp-2/pkg/errs"
	"fp-2/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService services.PhotoService
}

func NewPhotoHandler(photoService services.PhotoService) *photoHandler {
	return &photoHandler{photoService: photoService}
}

func (p *photoHandler) CreatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &entities.User{}
	user.ID = userId
	user.Username = username

	var requestBody dto.CreatePhotoRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := requestBody.ValidateStruct(); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response, err := p.photoService.CreatePhoto(user, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (p *photoHandler) GetAllPhotos(ctx *gin.Context) {
	response, err := p.photoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *photoHandler) UpdatePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		idError := errs.NewBadRequest(err.Error())
		ctx.JSON(idError.StatusCode(), idError)
		return
	}
	if photoId < 0 {
		idError := errs.NewBadRequest("Photo ID value must be positive")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.UpdatePhotoRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := requestBody.ValidateStruct(); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	updatedPhoto, messageErr := p.photoService.UpdatePhoto(uint(photoId), &requestBody)
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p *photoHandler) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		idError := errs.NewBadRequest(err.Error())
		ctx.JSON(idError.StatusCode(), idError)
		return
	}
	if photoId < 0 {
		idError := errs.NewBadRequest("Photo ID value must be positive")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	response, messageErr := p.photoService.DeletePhoto(uint(photoId))
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
