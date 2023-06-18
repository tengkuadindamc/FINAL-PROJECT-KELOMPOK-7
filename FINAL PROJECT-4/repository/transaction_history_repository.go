package repository

import (
	"finalproject4/model"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	GetAllTransactionHistory(transaction *[]model.TransactionHistory) error
	GetTransactionHistoryByUserId(transaction *[]model.TransactionHistory) error
	GetTransactionHistoryByID(id int) (model.TransactionHistory, error)
	CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error)
	DeleteTransactionHistory(transaction model.TransactionHistory) error
	GetUserByID(id int) (model.User, error)
	UpdateBalance(user model.User) (model.User, error)
	FindById(id int) (model.Product, error)
	EditProduct(product model.Product) error
	GetCategoryByID(id int) (model.Category, error)
	UpdateCategory(id int, category model.Category) (model.Category, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{db}
}

func (r *transactionHistoryRepository) GetAllTransactionHistory(transactionHistory *[]model.TransactionHistory) error {
	err := r.db.Preload("Product").Preload("User").Find(&transactionHistory).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionHistoryRepository) GetTransactionHistoryByUserId(transaction *[]model.TransactionHistory) error {
	err := r.db.Preload("Product").Find(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionHistoryRepository) GetTransactionHistoryByID(id int) (model.TransactionHistory, error) {
	var transactionHistory model.TransactionHistory
	err := r.db.Preload("Product").Preload("User").Where("id = ?", id).Find(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}
	return transactionHistory, nil
}

func (r *transactionHistoryRepository) CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error) {
	err := r.db.Create(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}
	return transactionHistory, nil
}

func (r *transactionHistoryRepository) DeleteTransactionHistory(transaction model.TransactionHistory) error {
	err := r.db.Debug().Delete(&transaction).Error
	if err == nil {
		return err
	}
	return err
}

func (r *transactionHistoryRepository) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *transactionHistoryRepository) UpdateBalance(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *transactionHistoryRepository) FindById(id int) (model.Product, error) {
	var product model.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *transactionHistoryRepository) EditProduct(product model.Product) error {
	err := r.db.Save(&product).Error
	return err
}

func (r *transactionHistoryRepository) GetCategoryByID(id int) (model.Category, error) {
	var category model.Category
	err := r.db.Preload("Products").Find(&category, id).Error
	return category, err
}

func (r *transactionHistoryRepository) UpdateCategory(id int, category model.Category) (model.Category, error) {
	err := r.db.Where("id = ?", id).Updates(&category).Error
	return category, err
}
