package services

import (
	"fp-2/dto"
	"fp-2/entities"
	"fp-2/pkg/errs"
	"fp-2/repositories"
)

type PhotoService interface {
	CreatePhoto(user *entities.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)
	GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr)
	UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo repositories.PhotoRepository
	userRepo  repositories.UserRepository
}

func NewPhotoService(photoRepo repositories.PhotoRepository, userRepo repositories.UserRepository) PhotoService {
	return &photoService{photoRepo: photoRepo, userRepo: userRepo}
}

func (p *photoService) CreatePhoto(user *entities.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	photo := payload.ToEntity()

	newPhoto, err := p.photoRepo.CreatePhoto(user, photo)
	if err != nil {
		return nil, err
	}

	response := &dto.CreatePhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserId,
		CreatedAt: newPhoto.CreatedAt,
	}

	return response, nil
}

func (p *photoService) GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr) {
	photos, err := p.photoRepo.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllPhotosResponse{}
	for _, photo := range photos {
		user := &entities.User{}
		user.ID = photo.UserId

		err := p.userRepo.GetUserByID(user)
		if err != nil {
			return nil, err
		}

		userOfPhotoResponse := dto.UserOfPhoto{
			Email:    user.Email,
			Username: user.Username,
		}

		response = append(response, dto.GetAllPhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserId,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      userOfPhotoResponse,
		})
	}

	return response, nil
}

func (p *photoService) UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
	photo, err := p.photoRepo.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}

	newPhoto := payload.ToEntity()

	updatedPhoto, messageErr := p.photoRepo.UpdatePhoto(photo, newPhoto)
	if messageErr != nil {
		return nil, messageErr
	}

	response := &dto.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoURL:  updatedPhoto.PhotoURL,
		UserID:    updatedPhoto.UserId,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr) {
	if err := p.photoRepo.DeletePhoto(id); err != nil {
		return nil, err
	}

	response := &dto.DeletePhotoResponse{
		Message: "photo has been successfully deleted",
	}

	return response, nil
}
