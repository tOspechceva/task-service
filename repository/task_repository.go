package repository

import (
	"database/sql"
	"task-service/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

// =====================
// CREATE
// =====================
func (r *TaskRepository) Create(task *models.Task) error {

	query := `
INSERT INTO tasks (
user_id, parent_task_id, title, description,
status_id, priority_id, due_date
)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id, created_at, updated_at
`

	return r.DB.QueryRow(
		query,
		task.UserID,
		task.ParentTaskID,
		task.Title,
		task.Description,
		task.StatusID,
		task.PriorityID,
		task.DueDate,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

// =====================
// GET BY ID
// =====================
func (r *TaskRepository) GetByID(id string) (*models.Task, error) {

	query := `
SELECT id,user_id,parent_task_id,title,description,
status_id,priority_id,due_date,completed_at,
is_completed,order_index,created_at,updated_at
FROM tasks WHERE id=$1
`

	row := r.DB.QueryRow(query, id)

	var task models.Task
	err := row.Scan(
		&task.ID,
		&task.UserID,
		&task.ParentTaskID,
		&task.Title,
		&task.Description,
		&task.StatusID,
		&task.PriorityID,
		&task.DueDate,
		&task.CompletedAt,
		&task.IsCompleted,
		&task.OrderIndex,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// =====================
// LIST BY USER
// =====================
func (r *TaskRepository) List(userID string) ([]models.Task, error) {

	query := `
SELECT id,user_id,parent_task_id,title,description,
status_id,priority_id,due_date,completed_at,
is_completed,order_index,created_at,updated_at
FROM tasks
WHERE user_id=$1
ORDER BY order_index
`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var t models.Task

		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.ParentTaskID,
			&t.Title,
			&t.Description,
			&t.StatusID,
			&t.PriorityID,
			&t.DueDate,
			&t.CompletedAt,
			&t.IsCompleted,
			&t.OrderIndex,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

// =====================
// UPDATE
// =====================
func (r *TaskRepository) Update(task *models.Task) error {

	query := `
UPDATE tasks SET
title=$1,
description=$2,
status_id=$3,
priority_id=$4,
due_date=$5,
is_completed=$6,
completed_at=$7,
updated_at=NOW()
WHERE id=$8
`

	_, err := r.DB.Exec(
		query,
		task.Title,
		task.Description,
		task.StatusID,
		task.PriorityID,
		task.DueDate,
		task.IsCompleted,
		task.CompletedAt,
		task.ID,
	)

	return err
}

// =====================
// DELETE
// =====================
func (r *TaskRepository) Delete(id string) error {

	_, err := r.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}

// =====================
// PAGINATION
// =====================
func (r *TaskRepository) ListWithPagination(
	userID string,
	limit int,
	offset int,
) ([]models.Task, error) {

	query := `
SELECT id,user_id,parent_task_id,title,description,
status_id,priority_id,due_date,completed_at,
is_completed,order_index,created_at,updated_at
FROM tasks
WHERE user_id=$1
ORDER BY order_index
LIMIT $2 OFFSET $3
`

	rows, err := r.DB.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.ParentTaskID,
			&t.Title,
			&t.Description,
			&t.StatusID,
			&t.PriorityID,
			&t.DueDate,
			&t.CompletedAt,
			&t.IsCompleted,
			&t.OrderIndex,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

