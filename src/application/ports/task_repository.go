// Package ports defines the interfaces (ports) that the application layer uses
// to communicate with external dependencies. This follows the Dependency Inversion Principle.
package ports

import (
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
)

// TaskRepository defines the contract for task persistence operations.
// Implementations of this interface are provided by the infrastructure layer.
type TaskRepository interface {
	Save(task *entities.Task) error
	FindById(id value_objects.TaskId) (*entities.Task, error)
	FindByStatus(status value_objects.TaskStatus) ([]*entities.Task, error)
	Delete(id value_objects.TaskId) error
}
