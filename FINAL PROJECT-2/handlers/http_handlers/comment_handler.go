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

type commentHandler struct {
	commentService services.CommentService
}

func NewCommentHandler(commentService services.CommentService) *commentHandler {
	return &commentHandler{commentService: commentService}
}

func (c *commentHandler) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &entities.User{}
	user.ID = userId
	user.Username = username

	var requestBody dto.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if messageErr := requestBody.ValidateStruct(); messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	createdComment, messageErr := c.commentService.CreateComment(user, &requestBody)
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusCreated, createdComment)
}

func (c *commentHandler) GetCommentsByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("currentUser").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &entities.User{
		Username: username,
	}
	user.ID = userId

	comments, messageErr := c.commentService.GetCommentsByUserId(userId)

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *commentHandler) GetCommentsByPhotoId(ctx *gin.Context) {
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

	comments, messageErr := c.commentService.GetCommentsByPhotoId(uint(photoId))
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *commentHandler) GetAllComment(ctx *gin.Context) {
	comments, messageErr := c.commentService.GetAllComment()
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *commentHandler) UpdateComment(ctx *gin.Context) {
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

	var updateRequest dto.UpdateCommentRequest
	err = ctx.ShouldBindJSON(&updateRequest)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	messageErr := updateRequest.ValidateStruct()
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	updateComment, messageErr := c.commentService.UpdateComment(uint(photoId), &updateRequest)
	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, updateComment)
}

func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		idError := errs.NewBadRequest(err.Error())
		ctx.JSON(idError.StatusCode(), idError)
		return
	}
	if commentId < 0 {
		idError := errs.NewBadRequest("Comment ID value must be positive")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	response, messageErr := c.commentService.DeleteComment(uint(commentId))

	if messageErr != nil {
		ctx.JSON(messageErr.StatusCode(), messageErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
