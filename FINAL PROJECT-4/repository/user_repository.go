package repository

import (
	"finalproject4/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user model.User) (model.User, error)
	Login(user model.User) (model.User, error)
	UpdateBalance(user model.User) (model.User, error)
	GetUserByID(id int) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) RegisterUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Login(user model.User) (model.User, error) {
	err := r.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateBalance(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
