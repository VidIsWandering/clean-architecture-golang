package usecases

import (
	"clean-architecture-golang/application/dto"
	"clean-architecture-golang/application/ports"
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
)

// GetTasksByStatusUseCase handles retrieving tasks filtered by status.
type GetTasksByStatusUseCase struct {
	Repo ports.TaskRepository
}

// Execute retrieves all tasks with the specified status.
// Returns a slice of TaskResponse DTOs.
func (uc *GetTasksByStatusUseCase) Execute(statusStr string) ([]dto.TaskResponse, error) {
	status := value_objects.TaskStatus(statusStr)
	if !status.IsValid() {
		return nil, entities.ErrInvalidStatus
	}
	tasks, err := uc.Repo.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	var responses []dto.TaskResponse
	for _, task := range tasks {
		responses = append(responses, dto.ToTaskResponse(task))
	}
	return responses, nil
}
