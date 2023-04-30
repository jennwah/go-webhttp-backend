package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/jennwah/go-webhttp-backend/domain"
	taskDomain "github.com/jennwah/go-webhttp-backend/domain/task"
)

type TaskController struct {
	TaskUsecase   taskDomain.TaskUsecase
	taskValidator *validator.Validate
}

func New(taskUseCase taskDomain.TaskUsecase) *TaskController {
	return &TaskController{
		taskUseCase,
		validator.New(),
	}
}

func (tc *TaskController) Create(c *gin.Context) {
	var task taskDomain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.taskValidator.Struct(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (tc *TaskController) GetByID(c *gin.Context) {
	taskID := c.Param("taskID")
	i, _ := strconv.ParseInt(taskID, 10, 64)

	task, err := tc.TaskUsecase.GetByID(c, i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
