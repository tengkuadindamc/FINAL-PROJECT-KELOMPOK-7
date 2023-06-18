package usecase

import (
	"errors"
	"final-project3/pkg/user/dto"
	"final-project3/pkg/user/model"
	"final-project3/pkg/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UsecaseInterfaceUser interface {
	Register(req dto.UserRequest) (*dto.UserResponse, error)
	Login(req dto.LoginRequest) (model.User, error)
	UpdateUserById(userId int, input model.User) (*dto.UserResponses, error)
	DeleteUserById(userId int) error
}

type usecaseUser struct {
	repository repository.RepositoryInterfaceUser
}

func InitUsecaseUser(repository repository.RepositoryInterfaceUser) UsecaseInterfaceUser {
	return &usecaseUser{
		repository: repository,
	}
}

// Register implements UsecaseInterfaceUser
func (u *usecaseUser) Register(req dto.UserRequest) (*dto.UserResponse, error) {
	isUserExist, _ := u.repository.GetUserByEmail(req.Email)

	if isUserExist.Id != 0 {
		return nil, errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	payload := model.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "member",
	}
	newUser, err := u.repository.CreateNewUser(payload)
	if err != nil {
		return nil, err
	}

	res := &dto.UserResponse{
		Id:        newUser.Id,
		Fullname:  newUser.Fullname,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}

	return res, nil
}

// Login implements UsecaseInterfaceUser
func (u *usecaseUser) Login(req dto.LoginRequest) (model.User, error) {
	user, err := u.repository.GetUserByEmail(req.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

// UpdateUserById implements UsecaseInterfaceUser
func (u *usecaseUser) UpdateUserById(userId int, input model.User) (*dto.UserResponses, error) {
	payload := model.User{
		Fullname: input.Fullname,
		Email:    input.Email,
	}
	user, err := u.repository.UpdateUserById(userId, payload)
	if err != nil {
		return nil, err
	} 

	res := &dto.UserResponses{
		Id:        input.Id,
		Fullname:  user.Fullname,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

// DeleleUserById implements UsecaseInterfaceUser
func (u *usecaseUser) DeleteUserById(userId int) error {
	err := u.repository.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return nil
}
