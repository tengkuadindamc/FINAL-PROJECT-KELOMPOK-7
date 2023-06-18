package dto

import (
	"fp-2/entities"
	"fp-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreatePhotoRequest struct {
	Title    string `json:"title" valid:"required~title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~photo url is required, url~Invalid photo url format"`
}

func (p *CreatePhotoRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (p *CreatePhotoRequest) ToEntity() *entities.Photo {
	return &entities.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type CreatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllPhotosResponse struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Caption   string      `json:"caption"`
	PhotoURL  string      `json:"photo_url"`
	UserID    uint        `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      UserOfPhoto `json:"user"`
}

type UserOfPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhotoRequest CreatePhotoRequest

func (p *UpdatePhotoRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (p *UpdatePhotoRequest) ToEntity() *entities.Photo {
	return &entities.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
