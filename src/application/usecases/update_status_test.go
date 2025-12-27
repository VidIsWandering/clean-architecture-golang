package usecases

import (
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
	"errors"
	"testing"
)

type mockRepoUpdate struct {
	found     *entities.Task
	findErr   error
	lastSaved *entities.Task
}

func (m *mockRepoUpdate) Save(task *entities.Task) error {
	m.lastSaved = task
	return nil
}
func (m *mockRepoUpdate) FindById(id value_objects.TaskId) (*entities.Task, error) {
	return m.found, m.findErr
}
func (m *mockRepoUpdate) FindByStatus(status value_objects.TaskStatus) ([]*entities.Task, error) {
	return nil, nil
}
func (m *mockRepoUpdate) Delete(id value_objects.TaskId) error { return nil }

func TestUpdateStatus_InvalidStatus(t *testing.T) {
	validId := string(value_objects.NewTaskId())
	repo := &mockRepoUpdate{found: &entities.Task{Status: value_objects.StatusTodo}}
	uc := &UpdateTaskStatusUseCase{Repo: repo}
	err := uc.Execute(validId, "invalid")
	if !errors.Is(err, entities.ErrInvalidStatus) {
		t.Errorf("Expected ErrInvalidStatus, got %v", err)
	}
}

func TestUpdateStatus_DoneToTodo(t *testing.T) {
	validId := string(value_objects.NewTaskId())
	repo := &mockRepoUpdate{found: &entities.Task{Status: value_objects.StatusDone}}
	uc := &UpdateTaskStatusUseCase{Repo: repo}
	err := uc.Execute(validId, "todo")
	if !errors.Is(err, entities.ErrInvalidTransition) {
		t.Errorf("Expected ErrInvalidTransition, got %v", err)
	}
}

func TestUpdateStatus_Success(t *testing.T) {
	validId := string(value_objects.NewTaskId())
	repo := &mockRepoUpdate{found: &entities.Task{Status: value_objects.StatusTodo}}
	uc := &UpdateTaskStatusUseCase{Repo: repo}
	err := uc.Execute(validId, "doing")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if repo.lastSaved.Status != value_objects.StatusDoing {
		t.Errorf("Expected status 'doing', got %v", repo.lastSaved.Status)
	}
}

func TestUpdateStatus_InvalidID(t *testing.T) {
	repo := &mockRepoUpdate{found: &entities.Task{Status: value_objects.StatusTodo}}
	uc := &UpdateTaskStatusUseCase{Repo: repo}
	err := uc.Execute("", "doing")
	if !errors.Is(err, ErrInvalidID) {
		t.Errorf("Expected ErrInvalidID, got %v", err)
	}
}
