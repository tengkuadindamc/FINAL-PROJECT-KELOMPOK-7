package services

import (
	"fp-2/dto"
	"fp-2/entities"
	"fp-2/pkg/errs"
	"fp-2/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(payload *dto.CreateSocialMediaRequest, userId uint) (*dto.CreateSocialMediaResponse, errs.MessageErr)
	GetAllSocialMedias() (*dto.AllSocialMediasResponse, errs.MessageErr)
	UpdateSocialMedia(sm_id uint, payload *dto.CreateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(sm_id uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepository
	userRepo        repositories.UserRepository
}

func NewSocialMediaService(socialMediaRepo repositories.SocialMediaRepository, userRepo repositories.UserRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepo: socialMediaRepo, userRepo: userRepo}
}

func (sm *socialMediaService) CreateSocialMedia(payload *dto.CreateSocialMediaRequest, userId uint) (*dto.CreateSocialMediaResponse, errs.MessageErr) {
	newSocialMedia := payload.ToEntity()
	newSocialMedia.UserId = userId

	createdSocialMedia, err := sm.socialMediaRepo.CreateSocialMedia(newSocialMedia)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateSocialMediaResponse{
		ID:             createdSocialMedia.ID,
		Name:           createdSocialMedia.Name,
		SocialMediaUrl: createdSocialMedia.SocialMediaUrl,
		UserId:         createdSocialMedia.UserId,
		CreatedAt:      createdSocialMedia.CreatedAt,
	}

	return response, nil
}

func (sm *socialMediaService) GetAllSocialMedias() (*dto.AllSocialMediasResponse, errs.MessageErr) {
	socialMedias, jumlahData, err := sm.socialMediaRepo.GetAllSocialMedias()
	if err != nil {
		return nil, err
	}

	var smListResponse []dto.SocialMediaResponse = make([]dto.SocialMediaResponse, 0, jumlahData)

	var userOfSMResponse dto.UserOfAnSMResponse
	var smResponse dto.SocialMediaResponse

	for _, smObject := range socialMedias {
		userOfSM := &entities.User{}
		userOfSM.ID = smObject.UserId

		err := sm.userRepo.GetUserByID(userOfSM)
		if err != nil {
			return nil, err
		}

		userOfSMResponse = dto.UserOfAnSMResponse{
			ID:       userOfSM.ID,
			Username: userOfSM.Username,
		}
		smResponse = dto.SocialMediaResponse{
			ID:             smObject.ID,
			Name:           smObject.Name,
			SocialMediaUrl: smObject.SocialMediaUrl,
			UserId:         smObject.UserId,
			CreatedAt:      smObject.CreatedAt,
			UpdatedAt:      smObject.UpdatedAt,
			User:           userOfSMResponse,
		}
		smListResponse = append(smListResponse, smResponse)
	}

	response := &dto.AllSocialMediasResponse{
		SocialMedias: smListResponse,
	}

	return response, nil
}

func (sm *socialMediaService) UpdateSocialMedia(smId uint, payload *dto.CreateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	socialMedia := payload.ToEntity()
	socialMedia.ID = smId

	updatedSocialMedia, err := sm.socialMediaRepo.UpdateSocialMedia(socialMedia)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateSocialMediaResponse{
		ID:             updatedSocialMedia.ID,
		Name:           updatedSocialMedia.Name,
		SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
		UserId:         updatedSocialMedia.UserId,
		UpdatedAt:      updatedSocialMedia.UpdatedAt,
	}

	return response, nil
}

func (sm *socialMediaService) DeleteSocialMedia(smId uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	err := sm.socialMediaRepo.DeleteSocialMedia(smId)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteSocialMediaResponse{
		Message: "Social media has been successfully deleted",
	}

	return response, nil
}
