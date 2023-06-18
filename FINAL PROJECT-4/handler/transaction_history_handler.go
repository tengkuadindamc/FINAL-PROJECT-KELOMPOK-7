package handler

import (
	"finalproject4/config"
	"finalproject4/model"
	"finalproject4/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHistoryHandler interface {
	GetAllTransactionHistory(ctx *gin.Context)
	GetTransactionHistoryByUserId(ctx *gin.Context)
	CreateTransactionHistory(ctx *gin.Context)
	DeleteTransactionHistory(ctx *gin.Context)
}

type transactionHistoryHandler struct {
	transactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionHistoryService service.TransactionHistoryService) *transactionHistoryHandler {
	return &transactionHistoryHandler{transactionHistoryService}
}

func (th *transactionHistoryHandler) GetAllTransactionHistory(ctx *gin.Context) {

	role := ctx.MustGet("currentUserRole").(string)
	currentUser := ctx.MustGet("currentUser").(model.User)
	if role == "customer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "You are not authorized to access this endpoint",
		})
		return
	}

	id := int(currentUser.ID)
	transactionHistory, err := th.transactionHistoryService.GetAllTransactionHistory(role, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatGetAdminTransaction(transactionHistory)
	ctx.JSON(http.StatusOK, formatter)
}

func (th *transactionHistoryHandler) GetTransactionHistoryByUserId(ctx *gin.Context) {

	role := ctx.MustGet("currentUserRole").(string)
	currentUser := ctx.MustGet("currentUser").(model.User)
	id := int(currentUser.ID)
	transactionHistory, err := th.transactionHistoryService.GetAllTransactionHistory(role, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	formatter := model.FormatGetUserTransaction(transactionHistory)
	ctx.JSON(http.StatusOK, formatter)

}

func (th *transactionHistoryHandler) CreateTransactionHistory(ctx *gin.Context) {
	var transaction model.TransactionHistoryInput
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a valid type of product and quantity"})
		return
	}

	currentUser := ctx.MustGet("currentUser").(model.User)
	userID := int(currentUser.ID)

	transactionResponse, err := th.transactionHistoryService.CreateTransaction(transaction, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactionResponse)
}

func ConvertTransactionBill(th model.TransactionHistory) (model.TransactionBill, error) {
	db := config.GetDB()
	Product := model.Product{}
	err := db.First(&Product, th.ProductID).Error
	if err != nil {
		return model.TransactionBill{}, err
	}
	transactionBill := model.TransactionBill{
		TotalPrice:   th.TotalPrice,
		Quantity:     th.Quantity,
		ProductTitle: Product.Title,
	}
	return transactionBill, err
}

func (h *transactionHistoryHandler) DeleteTransactionHistory(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	currentUser := ctx.MustGet("currentUser").(model.User)
	currentUserID := int(currentUser.ID)
	err = h.transactionHistoryService.DeleteTransactionHistory(id, currentUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	format := model.FormatDeleteTransaction()
	ctx.JSON(http.StatusOK, format)
}
