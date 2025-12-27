// Package repositories contains implementations of repository interfaces.
// These implementations handle data persistence and mapping between domain and persistence layers.
package repositories

import (
	"clean-architecture-golang/application/ports"
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
	"clean-architecture-golang/infrastructure/persistence"
	"errors"
	"sync"
)

// Sentinel error for not found
var ErrNotFound = errors.New("task not found")

// InMemoryTaskRepository implements ports.TaskRepository.
// It provides an in-memory implementation for task persistence.
type InMemoryTaskRepository struct {
	tasks map[string]*persistence.TaskModel
	mutex sync.RWMutex
}

// Ensure InMemoryTaskRepository implements ports.TaskRepository at compile time.
var _ ports.TaskRepository = (*InMemoryTaskRepository)(nil)

// NewInMemoryTaskRepository creates a new instance of InMemoryTaskRepository.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*persistence.TaskModel),
	}
}

// Save persists a task entity by converting it to a model and storing it.
func (r *InMemoryTaskRepository) Save(task *entities.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	model := persistence.FromDomain(task)
	r.tasks[string(task.ID)] = model
	return nil
}

// FindById retrieves a task by ID, converting the model back to a domain entity.
func (r *InMemoryTaskRepository) FindById(id value_objects.TaskId) (*entities.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	model, exists := r.tasks[string(id)]
	if !exists {
		return nil, ErrNotFound
	}
	return model.ToDomain(), nil
}

// FindByStatus retrieves all tasks with a specific status.
func (r *InMemoryTaskRepository) FindByStatus(status value_objects.TaskStatus) ([]*entities.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var tasks []*entities.Task
	for _, model := range r.tasks {
		if model.Status == status.String() {
			tasks = append(tasks, model.ToDomain())
		}
	}
	return tasks, nil
}

// Delete removes a task by ID.
func (r *InMemoryTaskRepository) Delete(id value_objects.TaskId) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.tasks[string(id)]; !exists {
		return ErrNotFound
	}
	delete(r.tasks, string(id))
	return nil
}
