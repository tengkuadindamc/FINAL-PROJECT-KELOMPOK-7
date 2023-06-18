package model

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string `json:"fullname" gorm:"type:varchar(30);not null;unique"`
	Email    string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password string `json:"password,omitempty" gorm:"size:255;not null"`
	Role     string `json:"role" gorm:"size:20;not null"`
	Balance  int    `json:"balance" gorm:"not null"`
}

type UserRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" bindig:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type UserBalanceRequest struct {
	Balance int `json:"balance"  binding:"required,min=0,max=500000000"`
}
type UserRegisterResponse struct {
	GormModel
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Passowrd string `json:"password"`
	Balance  int    `json:"balance"`
}

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = HashPass(u.Password)
	err = nil
	return
}
