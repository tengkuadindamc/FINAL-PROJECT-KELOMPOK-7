package controller

import (
	"final-project3/pkg/task/dto"
	"final-project3/pkg/task/usecase"
	"final-project3/utils/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TaskHTTPController struct {
	usecase usecase.UsecaseInterfaceTask
}

func InitControllerTask(uc usecase.UsecaseInterfaceTask) *TaskHTTPController {
	return &TaskHTTPController{
		usecase: uc,
	}
}

func (uc *TaskHTTPController) CreateNewTask(c *gin.Context) {
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}
	userInfo := c.MustGet("user_info").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	newReq := dto.TaskRequest{
		Title:       req.Title,
		Description: req.Description,
		Status:      false,
		UserId:      userId,
		CategoryId:  req.CategoryId,
	}
	task, err := uc.usecase.CreateNewTask(newReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (uc *TaskHTTPController) GetAllTask(c *gin.Context) {
	tasks, err := uc.usecase.GetAllTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (uc *TaskHTTPController) UpdateTaskById(c *gin.Context) {
	idString := c.Param("taskId")
	taskId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Task ID", "status": http.StatusBadRequest})
		return
	}
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	task, err := uc.usecase.UpdateTaskById(taskId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (uc *TaskHTTPController) UpdateStatusByTaskId(c *gin.Context) {
	idString := c.Param("taskId")
	taskId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Task ID", "status": http.StatusBadRequest})
		return
	}
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
		})
		return
	}

	task, err := uc.usecase.UpdateStatusByTaskId(taskId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (uc *TaskHTTPController) UpdateCategoryByTaskId(c *gin.Context) {
	idString := c.Param("taskId")
	taskId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Task ID", "status": http.StatusBadRequest})
		return
	}
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
		})
		return
	}

	task, err := uc.usecase.UpdateCategoryByTaskId(taskId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (uc *TaskHTTPController) DeleteTaskById(c *gin.Context) {
	idString := c.Param("taskId")
	taskId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Task ID", "status": http.StatusBadRequest})
		return
	}
	err = uc.usecase.DeleteTaskById(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task has been successfully deleted",
	})
}
