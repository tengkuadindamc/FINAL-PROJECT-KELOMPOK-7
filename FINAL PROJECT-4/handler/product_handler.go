package handler

import (
	"finalproject4/model"
	"finalproject4/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	EditProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	ViewProduct(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (p *productHandler) AddProduct(ctx *gin.Context) {
	var product model.AddProduct

	role_user := ctx.MustGet("currentUserRole").(string)

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	productData, err := p.productService.AddProduct(role_user, product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	productResponse := model.ResponseAddProduct{
		ID:         int(productData.ID),
		Title:      productData.Title,
		Price:      productData.Price,
		Stock:      productData.Stock,
		CategoryID: productData.CategoryID,
		CreatedAt:  productData.CreatedAt,
	}
	ctx.JSON(201, productResponse)
}

func (p *productHandler) EditProduct(ctx *gin.Context) {
	var editProduct model.EditProduct

	role_user := ctx.MustGet("currentUserRole").(string)

	err := ctx.ShouldBindJSON(&editProduct)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	productData, err := p.productService.EditProduct(id, role_user, editProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	productResponse := model.ResponseEditProduct{
		ID:         int(productData.ID),
		Title:      productData.Title,
		Price:      productData.Price,
		Stock:      productData.Stock,
		CategoryID: productData.CategoryID,
		CreatedAt:  productData.CreatedAt,
		UpdatedAt:  productData.UpdatedAt,
	}
	ctx.JSON(201, productResponse)
}

func (p *productHandler) DeleteProduct(ctx *gin.Context) {
	role_user := ctx.MustGet("currentUserRole").(string)
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	err := p.productService.DeleteProduct(role_user, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You're not allowed to do this!",
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted!",
	})
}

func (p *productHandler) ViewProduct(ctx *gin.Context) {
	product, err := p.productService.ViewProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := []model.Product(product)
	ctx.JSON(http.StatusOK, formatter)
}
