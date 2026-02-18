package services

import (
	"time"

	"task-service/models"
	"task-service/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

// CREATE
func (s *TaskService) Create(task *models.Task) error {
	return s.Repo.Create(task)
}

// GET
func (s *TaskService) Get(id string) (*models.Task, error) {
	return s.Repo.GetByID(id)
}

// LIST
func (s *TaskService) List(userID string) ([]models.Task, error) {
	return s.Repo.List(userID)
}

// UPDATE
func (s *TaskService) Update(task *models.Task) error {

	if task.IsCompleted && task.CompletedAt == nil {
		now := time.Now()
		task.CompletedAt = &now
	}

	return s.Repo.Update(task)
}

// DELETE
func (s *TaskService) Delete(id string) error {
	return s.Repo.Delete(id)
}
