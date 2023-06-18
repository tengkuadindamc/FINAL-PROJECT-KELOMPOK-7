package dto

import (
	"fp-2/entities"
	"fp-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~social media url is required, url~Invalid url format"`
}

func (sm *CreateSocialMediaRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(sm)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (sm *CreateSocialMediaRequest) ToEntity() *entities.SocialMedia {
	return &entities.SocialMedia{
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
	}
}

type CreateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserOfAnSMResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
type SocialMediaResponse struct {
	ID             uint               `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"social_media_url"`
	UserId         uint               `json:"user_id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	User           UserOfAnSMResponse `json:"user"`
}
type AllSocialMediasResponse struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}

type UpdateSocialMediaRequest CreateSocialMediaRequest

func (sm *UpdateSocialMediaRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(sm)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (sm *UpdateSocialMediaRequest) ToEntity() *entities.SocialMedia {
	return &entities.SocialMedia{
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
	}
}

type UpdateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DeleteSocialMediaResponse struct {
	Message string `json:"message"`
}
