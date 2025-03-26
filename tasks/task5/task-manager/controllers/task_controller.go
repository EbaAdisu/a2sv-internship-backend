package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
    TaskService *data.TaskService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
    return &TaskController{TaskService: taskService}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
    tasks := tc.TaskService.GetAllTasks()
    c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
    task, err := tc.TaskService.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	// println(task)
    if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
	// println(task)
    task = tc.TaskService.CreateTask(task)
    c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    updatedTask, err := tc.TaskService.UpdateTask(id, task)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
    if err := tc.TaskService.DeleteTask(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}