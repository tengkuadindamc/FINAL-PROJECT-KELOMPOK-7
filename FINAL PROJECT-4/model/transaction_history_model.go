package model

import (
	"time"
)

type TransactionHistory struct {
	GormModel
	UserID     int     `json:"user_id" gorm:"not null"`
	ProductID  int     `json:"product_id" gorm:"not null"`
	Quantity   int     `json:"quantity" gorm:"not null"`
	TotalPrice int     `json:"total_price" gorm:"not null"`
	Product    Product `gorm:"foreignKey:ProductID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User       User    `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransactionHistoryInput struct {
	ProductID int `json:"product_id" gorm:"not null"`
	Quantity  int `json:"quantity" gorm:"not null"`
}

type CategoryIdUri struct {
	ID int `uri:"id" binding:"required"`
}

type TransactionDeleteResponse struct {
	Message string `json:"message"`
}

type GetUserTransactionHistory struct {
	TransactionHistory []struct {
		ID         int `json:"id"`
		ProductID  int `json:"product_id"`
		UserID     int `json:"user_id"`
		Quantity   int `json:"quantity"`
		TotalPrice int `json:"total_price"`
		Product    struct {
			ID         int       `json:"id"`
			Title      string    `json:"title"`
			Price      int       `json:"price"`
			Stock      int       `json:"stock"`
			CategoryID int       `json:"category_id"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"product"`
	} `json:"transaction_history"`
}

type GetAdminTransactionHistory struct {
	TransactionHistory []struct {
		ID         int `json:"id"`
		ProductID  int `json:"product_id"`
		UserID     int `json:"user_id"`
		Quantity   int `json:"quantity"`
		TotalPrice int `json:"total_price"`
		Product    struct {
			ID         int       `json:"id"`
			Title      string    `json:"title"`
			Price      int       `json:"price"`
			Stock      int       `json:"stock"`
			CategoryID int       `json:"category_id"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"product"`
		User struct {
			ID        int       `json:"id"`
			Fullname  string    `json:"fullname"`
			Email     string    `json:"email"`
			Balance   int       `json:"balance"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"user"`
	} `json:"transaction_history"`
}

type CreateTransactionHistoryResponse struct {
	Message         string            `json:"message"`
	TransactionBill []TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}
