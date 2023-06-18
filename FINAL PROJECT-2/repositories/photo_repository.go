package repositories

import (
	"fp-2/entities"
	"fp-2/pkg/errs"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(user *entities.User, photo *entities.Photo) (*entities.Photo, errs.MessageErr)
	GetAllPhotos() ([]entities.Photo, errs.MessageErr)
	GetPhotoByID(id uint) (*entities.Photo, errs.MessageErr)
	UpdatePhoto(photo *entities.Photo, updatedPhoto *entities.Photo) (*entities.Photo, errs.MessageErr)
	DeletePhoto(id uint) errs.MessageErr
}

type photoPg struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) PhotoRepository {
	return &photoPg{db: db}
}

func (p *photoPg) CreatePhoto(user *entities.User, photo *entities.Photo) (*entities.Photo, errs.MessageErr) {
	photo.UserId = user.ID

	if err := p.db.Create(photo).Error; err != nil {
		return nil, errs.NewInternalServerError("can't create photo")
	}

	return photo, nil
}

func (p *photoPg) GetAllPhotos() ([]entities.Photo, errs.MessageErr) {
	var allPhotos []entities.Photo
	if err := p.db.Find(&allPhotos).Error; err != nil {
		return nil, errs.NewInternalServerError("can't get all photos")
	}

	return allPhotos, nil
}

func (p *photoPg) GetPhotoByID(id uint) (*entities.Photo, errs.MessageErr) {
	photo := entities.Photo{}
	if err := p.db.First(&photo, id).Error; err != nil {
		return nil, errs.NewNotFound("photo not found")
	}

	return &photo, nil
}

func (p *photoPg) UpdatePhoto(photo *entities.Photo, updatedPhoto *entities.Photo) (*entities.Photo, errs.MessageErr) {
	if err := p.db.Model(photo).Updates(updatedPhoto).Error; err != nil {
		return nil, errs.NewInternalServerError("can't update this photo")
	}

	return photo, nil
}

func (p *photoPg) DeletePhoto(id uint) errs.MessageErr {
	if err := p.db.Delete(&entities.Photo{}, id).Error; err != nil {
		return errs.NewInternalServerError("can't delete this photo")
	}

	return nil
}
