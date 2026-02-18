package handlers

import (
	"database/sql"
	"fmt"
	//"net/http"

	"task-service/utils"
	"task-service/dto"
	"task-service/models"
	"task-service/repository"
	"task-service/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Service *services.TaskService
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	repo := repository.NewTaskRepository(db)
	service := services.NewTaskService(repo)
	return &TaskHandler{Service: service}
}

// ======================
// CREATE
// ======================
func (h *TaskHandler) Create(c *gin.Context) {

	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false})
		return
	}

	userID := c.GetHeader("X-User-Id") // из gateway

	task := models.Task{
		UserID:       userID,
		ParentTaskID: req.ParentTaskID,
		Title:        req.Title,
		Description:  req.Description,
		StatusID:     req.StatusID,
		PriorityID:   req.PriorityID,
		DueDate:      req.DueDate,
	}

	err := h.Service.Create(&task)
	if err != nil {
		utils.Error(c, 500, err)
		return
	}

	c.JSON(201, gin.H{"success": true, "data": task})
}

// ======================
// GET
// ======================
func (h *TaskHandler) Get(c *gin.Context) {

	id := c.Param("id")

	task, err := h.Service.Get(id)
	if err != nil {
		utils.Error(c, 404, err)
		return
	}

	c.JSON(200, gin.H{"success": true, "data": task})
}

// ======================
// LIST
// ======================
func (h *TaskHandler) List(c *gin.Context) {

	userID := c.GetHeader("X-User-Id")

	page := 1
	limit := 20

	if p := c.Query("page"); p != "" {
		fmt.Sscan(p, &page)
	}

	if l := c.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}

	offset := (page - 1) * limit

	tasks, err := h.Service.Repo.ListWithPagination(userID, limit, offset)
	if err != nil {
		utils.Error(c, 500, err)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": tasks,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}

// ======================
// UPDATE
// ======================
func (h *TaskHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false})
		return
	}

	task, err := h.Service.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"success": false})
		return
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = req.Description
	}
	if req.StatusID != nil {
		task.StatusID = *req.StatusID
	}
	if req.PriorityID != nil {
		task.PriorityID = *req.PriorityID
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	if req.IsCompleted != nil {
		task.IsCompleted = *req.IsCompleted
	}

	err = h.Service.Update(task)
	if err != nil {
		utils.Error(c, 400, err)
		return
	}

	c.JSON(200, gin.H{"success": true})
}

// ======================
// DELETE
// ======================
func (h *TaskHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := h.Service.Delete(id)
	if err != nil {
		utils.Error(c, 500, err)
		return
	}

	c.JSON(200, gin.H{"success": true})
}
