package repository

import (
	"final-project3/pkg/user/model"

	"gorm.io/gorm"
)

type RepositoryInterfaceUser interface {
	CreateNewUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	UpdateUserById(userId int, user model.User) (model.User, error)
	DeleteUserById(userId int) error
}

type repositoryUser struct {
	db *gorm.DB
}

func InitRepositoryUser(db *gorm.DB) RepositoryInterfaceUser {
	db.AutoMigrate(&model.User{})
	return &repositoryUser{
		db: db,
	}
}

// CreateNewUser implements RepositoryInterfaceUser
func (r *repositoryUser) CreateNewUser(user model.User) (model.User, error) {
	if err := r.db.Table("users").Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByEmail implements RepositoryInterfaceUser
func (r *repositoryUser) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUserById implements RepositoryInterfaceUser
func (r *repositoryUser) UpdateUserById(userId int, user model.User) (model.User, error) {
	if err := r.db.Table("users").Where("id = ?", userId).Updates(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// DeleleUserById implements RepositoryInterfaceUser
func (r *repositoryUser) DeleteUserById(userId int) error {
	if err := r.db.Table("users").Where("id = ?", userId).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
