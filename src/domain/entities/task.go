// Package entities contains the core business entities of the domain layer.
// These entities encapsulate business rules and are independent of any external frameworks.
package entities

import (
	"clean-architecture-golang/domain/value_objects"
	"errors"
	"fmt"
	"time"
)

// ErrInvalidInput is a wrapper for validation errors.
var ErrInvalidInput = errors.New("invalid input")

// Sentinel errors for domain business rules
var (
	ErrEmptyTitle        = fmt.Errorf("%w: title cannot be empty", ErrInvalidInput)
	ErrInvalidStatus     = fmt.Errorf("%w: invalid status", ErrInvalidInput)
	ErrInvalidTransition = fmt.Errorf("%w: cannot change status from done to todo", ErrInvalidInput)
)

// Task represents a personal task with its core attributes and business rules.
type Task struct {
	ID          value_objects.TaskId
	Title       string
	Description string
	Status      value_objects.TaskStatus
	CreatedAt   time.Time
}

// NewTask creates a new task with validation.
// It enforces the business rule that title cannot be empty.
// Returns an error if validation fails.
func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return &Task{
		ID:          value_objects.NewTaskId(),
		Title:       title,
		Description: description,
		Status:      value_objects.StatusTodo,
		CreatedAt:   time.Now(),
	}, nil
}

// UpdateStatus changes the task status with business rule validation.
// Prevents invalid status transitions (e.g., DONE to TODO).
// Returns an error if the status is invalid or transition is not allowed.
func (t *Task) UpdateStatus(newStatus value_objects.TaskStatus) error {
	if !newStatus.IsValid() {
		return ErrInvalidStatus
	}
	if t.Status == value_objects.StatusDone && newStatus == value_objects.StatusTodo {
		return ErrInvalidTransition
	}
	t.Status = newStatus
	return nil
}
