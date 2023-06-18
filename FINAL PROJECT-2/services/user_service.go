package services

import (
	"fp-2/dto"
	"fp-2/helpers"
	"fp-2/pkg/errs"
	"fp-2/repositories"
)

type UserService interface {
	RegisterUser(payload *dto.CreateUserRequest) (*dto.CreateUserResponse, errs.MessageErr)
	LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	UpdateUser(id uint, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) RegisterUser(payload *dto.CreateUserRequest) (*dto.CreateUserResponse, errs.MessageErr) {
	user := payload.ToEntity()

	createdUser, err := u.userRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateUserResponse{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		ID:       createdUser.ID,
		Age:      createdUser.Age,
	}

	return response, nil
}

func (u *userService) LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr) {
	user := payload.ToEntity()
	passwordFromRequest := user.Password

	err := u.userRepo.LoginUser(user)

	if err != nil {
		return nil, err
	}

	ok := helpers.ComparePass([]byte(user.Password), []byte(passwordFromRequest))

	if !ok {
		return nil, errs.NewBadRequest("mismatch username and password, or the account is not found")
	}

	token := helpers.GenerateToken(user.ID, user.Username)

	response := &dto.LoginUserResponse{Token: token}

	return response, nil
}

func (u *userService) UpdateUser(id uint, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	user := payload.ToEntity()
	user.ID = id

	updatedUser, err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		ID:        updatedUser.ID,
		Age:       updatedUser.Age,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{Message: "account has been successfully deleted"}

	return response, nil
}
