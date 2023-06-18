package repositories

import (
	"fp-2/entities"
	"fp-2/pkg/errs"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(sm *entities.SocialMedia) (*entities.SocialMedia, errs.MessageErr)
	GetAllSocialMedias() ([]entities.SocialMedia, uint, errs.MessageErr)
	UpdateSocialMedia(sm *entities.SocialMedia) (*entities.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(smId uint) errs.MessageErr
}

type socialMediaPG struct {
	db *gorm.DB
}

func NewSocialMediaPG(db *gorm.DB) SocialMediaRepository {
	return &socialMediaPG{db: db}
}

func (sm *socialMediaPG) CreateSocialMedia(newSocialMedia *entities.SocialMedia) (*entities.SocialMedia, errs.MessageErr) {
	if err := sm.db.Create(newSocialMedia).Error; err != nil {
		return nil, errs.NewInternalServerError("can't create this social media")
	}

	return newSocialMedia, nil
}

func (sm *socialMediaPG) GetAllSocialMedias() ([]entities.SocialMedia, uint, errs.MessageErr) {
	var allSocialMedias []entities.SocialMedia
	result := sm.db.Find(&allSocialMedias)

	if err := result.Error; err != nil {
		return nil, 0, errs.NewInternalServerError("can't get all social medias")
	}

	jumlahData := result.RowsAffected

	return allSocialMedias, uint(jumlahData), nil
}

func (sm *socialMediaPG) UpdateSocialMedia(updatedSocialMedia *entities.SocialMedia) (*entities.SocialMedia, errs.MessageErr) {
	socialMedia := &entities.SocialMedia{}
	err := sm.db.Where("id = ?", updatedSocialMedia.ID).Take(&socialMedia).Error

	if err != nil {
		return nil, errs.NewNotFound("social media not found")
	}

	updatedSocialMedia.UserId = socialMedia.UserId
	err = sm.db.Model(updatedSocialMedia).Updates(updatedSocialMedia).Error

	if err != nil {
		return nil, errs.NewBadRequest(err.Error())
	}

	return updatedSocialMedia, nil
}

func (sm *socialMediaPG) DeleteSocialMedia(smId uint) errs.MessageErr {
	deleteSocialMedia := &entities.SocialMedia{}
	deleteSocialMedia.ID = smId

	err := sm.db.Delete(deleteSocialMedia).Error

	if err != nil {
		messageErr := errs.NewInternalServerError(err.Error())
		return messageErr
	}

	return nil
}
