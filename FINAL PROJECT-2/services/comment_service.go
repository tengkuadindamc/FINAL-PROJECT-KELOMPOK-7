package services

import (
	"fp-2/dto"
	"fp-2/entities"
	"fp-2/pkg/errs"
	"fp-2/repositories"
)

type CommentService interface {
	CreateComment(user *entities.User, payload *dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr)
	GetAllComment() ([]dto.GetAllCommentResponse, errs.MessageErr)
	GetCommentsByUserId(userId uint) ([]dto.GetAllCommentResponse, errs.MessageErr)
	GetCommentsByPhotoId(photoId uint) ([]dto.GetAllCommentResponse, errs.MessageErr)
	UpdateComment(id uint, comment *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)
	DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepo repositories.CommentRepository
	photoRepo   repositories.PhotoRepository
	userRepo    repositories.UserRepository
}

func NewCommentService(
	commentRepo repositories.CommentRepository,
	photoRepo repositories.PhotoRepository,
	userRepo repositories.UserRepository,
) CommentService {
	return &commentService{
		commentRepo: commentRepo, photoRepo: photoRepo, userRepo: userRepo}
}

func (c *commentService) CreateComment(user *entities.User, payload *dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr) {
	comment := payload.ToEntity()

	_, errPhoto := c.photoRepo.GetPhotoByID(uint(comment.PhotoId))

	if errPhoto != nil {
		return nil, errPhoto
	}

	createdComment, err := c.commentRepo.CreateComment(user, comment)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateCommentResponse{
		ID:        createdComment.ID,
		Message:   createdComment.Message,
		PhotoId:   createdComment.PhotoId,
		UserId:    createdComment.UserId,
		CreatedAt: createdComment.CreatedAt,
	}
	return response, nil
}

func entitiesToGetAllCommentsResponse(comment entities.Comment, user *entities.User, photo *entities.Photo) dto.GetAllCommentResponse {
	userResponse := dto.UserOfCommentResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	photoResponse := dto.PhotoOfCommentResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserId:   photo.UserId,
	}
	response := dto.GetAllCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
		User:      userResponse,
		Photo:     photoResponse,
	}

	return response
}

func (c *commentService) GetCommentsByUserId(userId uint) ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetCommentsByUserId(userId)

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user := &entities.User{}
		user.ID = comment.UserId
		err := c.userRepo.GetUserByID(user)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := entitiesToGetAllCommentsResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) GetCommentsByPhotoId(photoId uint) ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetCommentsByPhotoId(photoId)
	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user := &entities.User{}
		user.ID = comment.UserId
		err := c.userRepo.GetUserByID(user)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := entitiesToGetAllCommentsResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) GetAllComment() ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user := &entities.User{}
		user.ID = comment.UserId
		err := c.userRepo.GetUserByID(user)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := entitiesToGetAllCommentsResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) UpdateComment(id uint, payload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	comment, err := c.commentRepo.GetCommentById(id)

	if err != nil {
		return nil, err
	}

	newComment := payload.ToEntity()

	updatedComment, err := c.commentRepo.UpdateComment(comment, newComment)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateCommentResponse{
		ID:        updatedComment.ID,
		Message:   updatedComment.Message,
		PhotoId:   updatedComment.PhotoId,
		UserId:    updatedComment.UserId,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr) {
	comment, err := c.commentRepo.GetCommentById(id)

	if err != nil {
		return nil, err
	}

	if err := c.commentRepo.DeleteComment(comment); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteCommentResponse{
		Message: "comment has been successfully deleted",
	}
	return deleteResponse, nil
}
