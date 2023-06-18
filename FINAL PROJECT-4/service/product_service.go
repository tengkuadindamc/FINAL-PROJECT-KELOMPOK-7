package service

import (
	"errors"
	"finalproject4/model"
	"finalproject4/repository"
)

type ProductService interface {
	AddProduct(role_user string, product model.AddProduct) (model.Product, error)
	EditProduct(id int, role_user string, product model.EditProduct) (model.Product, error)
	DeleteProduct(role_user string, id int) error
	ViewProduct() ([]model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (p *productService) AddProduct(role_user string, product model.AddProduct) (model.Product, error) {
	if role_user != "admin" {
		return model.Product{}, errors.New("You aren't allowed to do this! You are not Admin!")
	}
	var newProduct = model.Product{
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}

	productResponse, err := p.productRepository.AddProduct(newProduct)
	if err != nil {
		return model.Product{}, err
	}

	return productResponse, nil
}

func (p *productService) EditProduct(id int, role_user string, product model.EditProduct) (model.Product, error) {
	if role_user != "admin" {
		return model.Product{}, errors.New("You aren't allowed to do this! You are not Admin!")
	}

	editedProduct, err := p.productRepository.FindById(id)
	if err != nil {
		return model.Product{}, err
	}

	editedProduct.Title = product.Title
	editedProduct.Price = product.Price
	editedProduct.Stock = product.Stock
	editedProduct.CategoryID = product.CategoryID

	result, err := p.productRepository.EditProduct(editedProduct)
	return result, err
}

func (p *productService) DeleteProduct(role_user string, id int) error {
	if role_user != "admin" {
		return errors.New("You aren't allowed to do this! You are not Admin!")
	}

	deletedProduct, err := p.productRepository.FindById(id)
	if err != nil {
		return errors.New("There's nothing!")
	}

	err = p.productRepository.DeleteProduct(deletedProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) ViewProduct() ([]model.Product, error) {
	product, err := p.productRepository.ViewProduct()
	if err != nil {
		return product, err
	}
	return product, nil
}
