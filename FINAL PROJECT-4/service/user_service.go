package service

import (
	"finalproject4/model"
	"finalproject4/repository"
)

type UserService interface {
	RegisterUser(userRequest model.UserRegisterRequest) (model.UserRegisterResponse, error)
	Login(userLogin model.UserLoginRequest) (model.User, error)
	TopUpBalance(userBalance model.UserBalanceRequest, userID int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(userRequest model.UserRegisterRequest) (model.UserRegisterResponse, error) {
	var userResponse model.UserRegisterResponse
	user := model.User{
		Fullname: userRequest.FullName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     "customer",
		Balance:  0,
	}

	user, err := us.userRepository.RegisterUser(user)
	if err != nil {
		return userResponse, err
	}

	userResponse = model.UserRegisterResponse{
		GormModel: model.GormModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
		},
		Fullname: user.Fullname,
		Email:    user.Email,
		Passowrd: user.Password,
		Balance:  user.Balance,
	}

	return userResponse, nil
}

func (us *userService) Login(userLogin model.UserLoginRequest) (model.User, error) {
	var user model.User
	user.Email = userLogin.Email
	user.Password = userLogin.Password
	user, err := us.userRepository.Login(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *userService) TopUpBalance(userBalance model.UserBalanceRequest, userID int) error {
	var user model.User

	user, err := us.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Balance += userBalance.Balance
	user, err = us.userRepository.UpdateBalance(user)
	if err != nil {
		return err
	}

	return nil
}
