package model

import "time"

// func FormatTransaction(transaction TransactionHistory) CreateTransactionHistoryResponse {
// 	formatter := CreateTransactionHistoryResponse{
// 		Message: "You have successfully bought this product",
// 		TransactionBill: struct {
// 			TotalPrice   int    `json:"total_price"`
// 			Quantity     int    `json:"quantity"`
// 			ProductTitle string `json:"product_title"`
// 		}{
// 			TotalPrice:   transaction.TotalPrice,
// 			Quantity:     transaction.Quantity,
// 			ProductTitle: transaction.Product.Title,
// 		},
// 	}
// 	return formatter
// }

func FormatGetUserTransaction(transaction []TransactionHistory) GetUserTransactionHistory {
	var formatter GetUserTransactionHistory
	for _, value := range transaction {
		formatter.TransactionHistory = append(formatter.TransactionHistory, struct {
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
		}{
			ID:         int(value.ID),
			ProductID:  value.ProductID,
			UserID:     value.UserID,
			Quantity:   value.Quantity,
			TotalPrice: value.TotalPrice,
			Product: struct {
				ID         int       `json:"id"`
				Title      string    `json:"title"`
				Price      int       `json:"price"`
				Stock      int       `json:"stock"`
				CategoryID int       `json:"category_id"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
			}{
				ID:         int(value.Product.ID),
				Title:      value.Product.Title,
				Price:      value.Product.Price,
				Stock:      value.Product.Stock,
				CategoryID: value.Product.CategoryID,
				CreatedAt:  value.Product.CreatedAt,
				UpdatedAt:  value.Product.UpdatedAt,
			},
		})
	}
	return formatter
}

func FormatGetAdminTransaction(transaction []TransactionHistory) GetAdminTransactionHistory {
	var formatter GetAdminTransactionHistory
	for _, value := range transaction {
		formatter.TransactionHistory = append(formatter.TransactionHistory, struct {
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
		}{
			ID:         int(value.ID),
			ProductID:  value.ProductID,
			UserID:     value.UserID,
			Quantity:   value.Quantity,
			TotalPrice: value.TotalPrice,
			Product: struct {
				ID         int       `json:"id"`
				Title      string    `json:"title"`
				Price      int       `json:"price"`
				Stock      int       `json:"stock"`
				CategoryID int       `json:"category_id"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
			}{
				ID:         int(value.Product.ID),
				Title:      value.Product.Title,
				Price:      value.Product.Price,
				Stock:      value.Product.Stock,
				CategoryID: value.Product.CategoryID,
				CreatedAt:  value.Product.CreatedAt,
				UpdatedAt:  value.Product.UpdatedAt,
			},
			User: struct {
				ID        int       `json:"id"`
				Fullname  string    `json:"fullname"`
				Email     string    `json:"email"`
				Balance   int       `json:"balance"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			}{
				ID:        int(value.User.ID),
				Fullname:  value.User.Fullname,
				Email:     value.User.Email,
				Balance:   value.User.Balance,
				CreatedAt: value.User.CreatedAt,
				UpdatedAt: value.User.UpdatedAt,
			},
		})
	}
	return formatter
}

func FormatDeleteTransaction() TransactionDeleteResponse {
	formatter := TransactionDeleteResponse{
		Message: "Transaction deleted successfully",
	}
	return formatter
}
