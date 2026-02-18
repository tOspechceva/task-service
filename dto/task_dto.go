package dto

import "time"

type CreateTaskRequest struct {
	Title        string     `json:"title" binding:"required"`
	Description  *string    `json:"description"`
	ParentTaskID *string    `json:"parent_task_id"`
	StatusID     string     `json:"status_id" binding:"required"`
	PriorityID   string     `json:"priority_id" binding:"required"`
	DueDate      *time.Time `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title        *string    `json:"title"`
	Description  *string    `json:"description"`
	StatusID     *string    `json:"status_id"`
	PriorityID   *string    `json:"priority_id"`
	DueDate      *time.Time `json:"due_date"`
	IsCompleted  *bool      `json:"is_completed"`
}