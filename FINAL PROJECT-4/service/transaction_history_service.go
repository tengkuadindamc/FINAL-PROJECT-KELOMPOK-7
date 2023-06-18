package service

import (
	"errors"
	"finalproject4/model"
	"finalproject4/repository"
)

type TransactionHistoryService interface {
	GetAllTransactionHistory(role string, UserID int) ([]model.TransactionHistory, error)
	CreateTransaction(transactionHistory model.TransactionHistoryInput, UserID int) (model.CreateTransactionHistoryResponse, error)
	DeleteTransactionHistory(id_task int, UserID int) error
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository}
}

func (s *transactionHistoryService) GetAllTransactionHistory(role string, UserID int) ([]model.TransactionHistory, error) {
	var transactionHistory []model.TransactionHistory
	if role == "admin" {
		err := s.transactionHistoryRepository.GetAllTransactionHistory(&transactionHistory)
		if len(transactionHistory) == 0 {
			return transactionHistory, errors.New("no transaction history")
		}
		if err != nil {
			return []model.TransactionHistory{}, err
		}
		return transactionHistory, nil
	}
	if role == "customer" {
		err := s.transactionHistoryRepository.GetAllTransactionHistory(&transactionHistory)
		if len(transactionHistory) == 0 {
			return transactionHistory, errors.New("no transaction history")
		}
		if err != nil {
			return []model.TransactionHistory{}, err
		}
		var filteredTransactionHistory []model.TransactionHistory
		for _, transaction := range transactionHistory {
			if transaction.UserID == UserID {
				filteredTransactionHistory = append(filteredTransactionHistory, transaction)
			}
		}
		return filteredTransactionHistory, nil
	}
	return []model.TransactionHistory{}, errors.New("Login first only user or admin role access this")
}

func (s *transactionHistoryService) CreateTransaction(transaction model.TransactionHistoryInput, UserID int) (model.CreateTransactionHistoryResponse, error) {

	CreatedTransaction := model.TransactionHistory{
		UserID:    UserID,
		ProductID: transaction.ProductID,
		Quantity:  transaction.Quantity,
	}

	product, err := s.transactionHistoryRepository.FindById(transaction.ProductID)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, errors.New("product not found")
	}

	user, err := s.transactionHistoryRepository.GetUserByID(UserID)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, errors.New("user not found")
	}

	categoryData, err := s.transactionHistoryRepository.GetCategoryByID(product.CategoryID)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, errors.New("category not found")
	}
	if categoryData.ID == 0 {
		return model.CreateTransactionHistoryResponse{}, errors.New("category not found")
	}

	total := product.Stock - transaction.Quantity
	price := product.Price * transaction.Quantity
	categoryData.Sold_product_amount += transaction.Quantity

	// Check Stock
	if total < 0 {
		return model.CreateTransactionHistoryResponse{}, errors.New("product's stock are not enough")
	}

	// Check balance
	if user.Balance < price {
		return model.CreateTransactionHistoryResponse{}, errors.New("you don't have enough balance")
	}

	// update Product
	product.Stock = total

	err = s.transactionHistoryRepository.EditProduct(product)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, err
	}

	// update balance
	user.Balance = user.Balance - price

	user, err = s.transactionHistoryRepository.UpdateBalance(user)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, err
	}

	// update Category
	updatedCategory, err := s.transactionHistoryRepository.UpdateCategory(product.CategoryID, categoryData)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, errors.New("category not found")
	}
	if updatedCategory.ID == 0 {
		return model.CreateTransactionHistoryResponse{}, errors.New("category not found")
	}

	//Create a new transaction
	CreatedTransaction.TotalPrice = price
	_, err = s.transactionHistoryRepository.CreateTransactionHistory(CreatedTransaction)
	if err != nil {
		return model.CreateTransactionHistoryResponse{}, err
	}
	transactionBill := model.TransactionBill{
		TotalPrice:   price,
		ProductTitle: product.Title,
		Quantity:     transaction.Quantity,
	}
	transactionResponse := model.CreateTransactionHistoryResponse{
		Message:         "Transaction Success",
		TransactionBill: []model.TransactionBill{transactionBill},
	}
	return transactionResponse, nil
}

func (s *transactionHistoryService) DeleteTransactionHistory(id_task int, UserID int) error {
	transaction, err := s.transactionHistoryRepository.GetTransactionHistoryByID(id_task)
	if err != nil {
		return err
	}
	if transaction.UserID != UserID {
		return errors.New("you are not the owner of this transaction history")
	}
	err = s.transactionHistoryRepository.DeleteTransactionHistory(transaction)

	if err != nil {
		return err
	}
	return nil
}
