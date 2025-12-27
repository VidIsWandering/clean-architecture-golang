package usecases

import (
	"clean-architecture-golang/application/ports"
	"clean-architecture-golang/domain/value_objects"
)

// DeleteTaskUseCase handles the deletion of tasks.
type DeleteTaskUseCase struct {
	Repo ports.TaskRepository
}

// Execute deletes a task by its string ID.
// Returns an error if the task is not found or deletion fails.
func (uc *DeleteTaskUseCase) Execute(idStr string) error {
	parsedId, err := value_objects.ParseTaskId(idStr)
	if err != nil {
		return ErrInvalidID
	}
	return uc.Repo.Delete(parsedId)
}
