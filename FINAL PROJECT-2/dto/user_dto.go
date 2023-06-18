package dto

import (
	"fp-2/entities"
	"fp-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateUserRequest struct {
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"required~email is required, email~Invalid email format"`
	Password string `json:"password" valid:"required~password is required, minstringlength(6)~Minimum length of Password is 6 characters!"`
	Age      int    `json:"age" valid:"required~age is required"`
}

func (u *CreateUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	if u.Age < 9 {
		return errs.NewUnprocessableEntity("Invalid Age value")
	}

	return nil
}

func (u *CreateUserRequest) ToEntity() *entities.User {
	return &entities.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
	}
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
}

type LoginUserRequest struct {
	Username string `json:"username" valid:"required~username is required"`
	Password string `json:"password" valid:"required~password is required"`
}

func (u *LoginUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *LoginUserRequest) ToEntity() *entities.User {
	return &entities.User{
		Username: u.Username,
		Password: u.Password,
	}
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"required~email is required, email~Invalid email format"`
}

func (u *UpdateUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *UpdateUserRequest) ToEntity() *entities.User {
	return &entities.User{
		Username: u.Username,
		Email:    u.Email,
	}
}

type UpdateUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ID        uint      `json:"id"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
