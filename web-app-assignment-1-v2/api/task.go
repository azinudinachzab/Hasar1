package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) TaskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task data"})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to store task"})
		return
	}

	c.JSON(http.StatusOK, newTask)
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	// Implement the UpdateTask API endpoint
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	// Implement the DeleteTask API endpoint
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	// Implement the GetTaskByID API endpoint
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	// Implement the GetTaskList API endpoint
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	// Implement the GetTaskListByCategory API endpoint
}
