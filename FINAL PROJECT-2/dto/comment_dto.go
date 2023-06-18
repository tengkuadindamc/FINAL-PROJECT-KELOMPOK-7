package dto

import (
	"fp-2/entities"
	"fp-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateCommentRequest struct {
	Message string `json:"message" valid:"required~message is required"`
	PhotoId uint   `json:"photo_id" valid:"required~photo ID is required"`
}

func (u *CreateCommentRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (c *CreateCommentRequest) ToEntity() *entities.Comment {
	return &entities.Comment{
		Message: c.Message,
		PhotoId: c.PhotoId,
	}
}

type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserOfCommentResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoOfCommentResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}

type GetAllCommentResponse struct {
	ID        uint                   `json:"id"`
	Message   string                 `json:"message"`
	PhotoId   uint                   `json:"photo_id"`
	UserId    uint                   `json:"user_id"`
	UpdatedAt time.Time              `json:"updated_at"`
	CreatedAt time.Time              `json:"created_at"`
	User      UserOfCommentResponse  `json:"User"`
	Photo     PhotoOfCommentResponse `json:"Photo"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Message is required"`
}

func (u *UpdateCommentRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (c *UpdateCommentRequest) ToEntity() *entities.Comment {
	return &entities.Comment{
		Message: c.Message,
	}
}

type UpdateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
