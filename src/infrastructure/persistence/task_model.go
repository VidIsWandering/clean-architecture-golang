// Package persistence contains data models for persistence layer.
// These models include database-specific annotations and mapping functions.
package persistence

import (
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
	"time"
)

// TaskModel represents the database schema for tasks.
// It includes JSON tags for serialization and can be extended with ORM tags.
type TaskModel struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToDomain converts a TaskModel to a domain Task entity.
func (m *TaskModel) ToDomain() *entities.Task {
	return &entities.Task{
		ID:          value_objects.TaskId(m.ID),
		Title:       m.Title,
		Description: m.Description,
		Status:      value_objects.TaskStatus(m.Status),
		CreatedAt:   m.CreatedAt,
	}
}

// FromDomain converts a domain Task entity to a TaskModel.
func FromDomain(task *entities.Task) *TaskModel {
	return &TaskModel{
		ID:          string(task.ID),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status.String(),
		CreatedAt:   task.CreatedAt,
	}
}
