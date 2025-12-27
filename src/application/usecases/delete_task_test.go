package usecases

import (
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
	"clean-architecture-golang/infrastructure/repositories"
	"errors"
	"testing"
)

type mockRepoDelete struct {
	deleteErr error
	deletedId value_objects.TaskId
}

func (m *mockRepoDelete) Save(task *entities.Task) error                           { return nil }
func (m *mockRepoDelete) FindById(id value_objects.TaskId) (*entities.Task, error) { return nil, nil }
func (m *mockRepoDelete) FindByStatus(status value_objects.TaskStatus) ([]*entities.Task, error) {
	return nil, nil
}
func (m *mockRepoDelete) Delete(id value_objects.TaskId) error {
	m.deletedId = id
	return m.deleteErr
}

func TestDeleteTask_NotFound(t *testing.T) {
	validId := string(value_objects.NewTaskId())
	repo := &mockRepoDelete{deleteErr: repositories.ErrNotFound}
	uc := &DeleteTaskUseCase{Repo: repo}
	err := uc.Execute(validId)
	if !errors.Is(err, repositories.ErrNotFound) {
		t.Errorf("Expected repositories.ErrNotFound, got %v", err)
	}
}

func TestDeleteTask_Success(t *testing.T) {
	validId := string(value_objects.NewTaskId())
	repo := &mockRepoDelete{}
	uc := &DeleteTaskUseCase{Repo: repo}
	err := uc.Execute(validId)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if repo.deletedId != value_objects.TaskId(validId) {
		t.Errorf("Expected deleted id '%s', got %v", validId, repo.deletedId)
	}
}
