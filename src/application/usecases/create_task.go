// Package usecases contains the application use cases.
// Each use case orchestrates business logic using domain entities and external ports.
package usecases

import (
	"clean-architecture-golang/application/dto"
	"clean-architecture-golang/application/ports"
	"clean-architecture-golang/domain/entities"
)

// CreateTaskUseCase handles the creation of new tasks.
type CreateTaskUseCase struct {
	Repo ports.TaskRepository
}

// Execute creates a new task and persists it.
// Returns the created task as a DTO or an error if creation fails.
func (uc *CreateTaskUseCase) Execute(req dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	task, err := entities.NewTask(req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	err = uc.Repo.Save(task)
	if err != nil {
		return nil, err
	}
	response := dto.ToTaskResponse(task)
	return &response, nil
}
