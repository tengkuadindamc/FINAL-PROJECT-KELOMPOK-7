package handler

import (
	"finalproject4/model"
	"finalproject4/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler interface {
	CreateCategory(ctx *gin.Context)
	GetAllCategory(ctx *gin.Context)
	PatchCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoriesHandler struct {
	categoriesService service.CategoriesService
}

func NewCategoriesHandler(categoriesService service.CategoriesService) *categoriesHandler {
	return &categoriesHandler{categoriesService}
}

func (h *categoriesHandler) CreateCategory(ctx *gin.Context) {
	var categories model.CategoryInput

	role_user := ctx.MustGet("currentUserRole").(string)

	err := ctx.ShouldBindJSON(&categories)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid type of category using the JSON object `type`"})
		return
	}

	categoriesData, err := h.categoriesService.CreateCategory(role_user, categories)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoriesResponse := model.CategoryPostResponse{
		ID:        categoriesData.ID,
		Type:      categoriesData.Type,
		CreatedAt: categoriesData.CreatedAt,
	}
	ctx.JSON(http.StatusOK, categoriesResponse)
}

func (h *categoriesHandler) GetAllCategory(ctx *gin.Context) {
	var (
		allProducts   []model.CategoryProduct
		allCategories []model.CategoryGetResponse
	)

	categoryData, err := h.categoriesService.GetAllCategory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, dataCategory := range categoryData {

		for _, dataProduct := range dataCategory.Products {
			Product := convertProductResponse(dataProduct)
			allProducts = append(allProducts, Product)
		}

		category := convertCategoryResponse(dataCategory, allProducts)

		allCategories = append(allCategories, category)
		allProducts = nil
	}

	ctx.JSON(http.StatusOK, allCategories)
}

func convertCategoryResponse(dataCategory model.Category, allProducts []model.CategoryProduct) model.CategoryGetResponse {
	return model.CategoryGetResponse{
		Id:                  dataCategory.ID,
		Type:                dataCategory.Type,
		Sold_product_amount: dataCategory.Sold_product_amount,
		CreatedAt:           dataCategory.CreatedAt,
		UpdatedAt:           dataCategory.UpdatedAt,
		Products:            allProducts,
	}
}

func convertProductResponse(p model.Product) model.CategoryProduct {
	return model.CategoryProduct{
		GormModel: p.GormModel,
		Title:     p.Title,
		Price:     p.Price,
		Stock:     p.Stock,
	}
}

func (h *categoriesHandler) PatchCategory(ctx *gin.Context) {
	var category model.CategoryInput

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	role_user := ctx.MustGet("currentUserRole").(string)

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid type of category using the JSON object `type`"})
		return
	}

	categoriesData, err := h.categoriesService.PatchCategory(role_user, id, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryResponse := model.CategoryPatchResponse{
		ID:                  categoriesData.ID,
		Type:                categoriesData.Type,
		Sold_product_amount: categoriesData.Sold_product_amount,
		UpdatedAt:           categoriesData.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, categoryResponse)
}

func (h *categoriesHandler) DeleteCategory(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	role_user := ctx.MustGet("currentUserRole").(string)

	err := h.categoriesService.DeleteCategory(role_user, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category has been successfully deleted"})
}
