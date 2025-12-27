package usecases

import (
	"clean-architecture-golang/application/ports"
	"clean-architecture-golang/domain/value_objects"
	"errors"
)

// Package-level errors for usecases
var ErrInvalidID = errors.New("invalid id")

// UpdateTaskStatusUseCase handles updating the status of existing tasks.
type UpdateTaskStatusUseCase struct {
	Repo ports.TaskRepository
}

// Execute updates the status of a task identified by its string ID.
// Validates the new status and enforces business rules.
// Returns an error if the task is not found, status is invalid, or transition is not allowed.
func (uc *UpdateTaskStatusUseCase) Execute(idStr string, statusStr string) error {
	parsedId, err := value_objects.ParseTaskId(idStr)
	if err != nil {
		return ErrInvalidID
	}
	task, err := uc.Repo.FindById(parsedId)
	if err != nil {
		return err
	}
	newStatus := value_objects.TaskStatus(statusStr)
	err = task.UpdateStatus(newStatus)
	if err != nil {
		return err
	}
	return uc.Repo.Save(task)
}
