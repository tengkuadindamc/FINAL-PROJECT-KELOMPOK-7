package repositories

import (
	"fp-2/entities"
	"fp-2/pkg/errs"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(user *entities.User) errs.MessageErr
	RegisterUser(user *entities.User) (*entities.User, errs.MessageErr)
	LoginUser(user *entities.User) errs.MessageErr
	UpdateUser(user *entities.User) (*entities.User, errs.MessageErr)
	DeleteUser(id uint) errs.MessageErr
}

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) UserRepository {
	return &userPG{db: db}
}

func (u *userPG) GetUserByID(user *entities.User) errs.MessageErr {
	err := u.db.Where("id = ?", user.ID).Take(&user).Error

	if err != nil {
		return errs.NewNotFound("user not found")
	}

	return nil
}

func (u *userPG) RegisterUser(newUser *entities.User) (*entities.User, errs.MessageErr) {
	if err := u.db.Create(newUser).Error; err != nil {
		error := errs.NewInternalServerError("can't register this user")
		return nil, error
	}

	return newUser, nil
}

func (u *userPG) LoginUser(user *entities.User) errs.MessageErr {
	err := u.db.Where("username = ?", user.Username).Take(&user).Error

	if err != nil {
		messageErr := errs.NewBadRequest("mismatch username and password, or the account is not found")
		return messageErr
	}

	return nil
}

func (u *userPG) UpdateUser(user *entities.User) (*entities.User, errs.MessageErr) {
	if messageErr := u.GetUserByID(user); messageErr != nil {
		return nil, messageErr
	}

	err := u.db.Model(user).Updates(user).Error

	if err != nil {
		messageErr := errs.NewBadRequest("can't update this user")
		return nil, messageErr
	}

	return user, nil
}

func (u *userPG) DeleteUser(id uint) errs.MessageErr {
	initialUser := &entities.User{}
	initialUser.ID = id

	if err := u.GetUserByID(initialUser); err != nil {
		return errs.NewNotFound("user not found")
	}

	if err := u.db.Model(initialUser).Delete(initialUser).Error; err != nil {
		return errs.NewInternalServerError("can't delete this user")
	}

	return nil
}
