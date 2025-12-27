package usecases

import (
	"clean-architecture-golang/application/dto"
	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
	"errors"
	"testing"
)

type mockRepoCreate struct {
	lastSaved *entities.Task
	SaveErr   error
}

func (m *mockRepoCreate) Save(task *entities.Task) error {
	m.lastSaved = task
	return m.SaveErr
}
func (m *mockRepoCreate) FindById(id value_objects.TaskId) (*entities.Task, error) { return nil, nil }
func (m *mockRepoCreate) FindByStatus(status value_objects.TaskStatus) ([]*entities.Task, error) {
	return nil, nil
}
func (m *mockRepoCreate) Delete(id value_objects.TaskId) error { return nil }

func TestCreateTask_EmptyTitle(t *testing.T) {
	repo := &mockRepoCreate{}
	uc := &CreateTaskUseCase{Repo: repo}
	_, err := uc.Execute(dto.CreateTaskRequest{Title: "", Description: "desc"})
	if !errors.Is(err, entities.ErrEmptyTitle) {
		t.Errorf("Expected ErrEmptyTitle, got %v", err)
	}
}

func TestCreateTask_Success(t *testing.T) {
	repo := &mockRepoCreate{}
	uc := &CreateTaskUseCase{Repo: repo}
	resp, err := uc.Execute(dto.CreateTaskRequest{Title: "abc", Description: "desc"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.Title != "abc" {
		t.Errorf("Expected title 'abc', got %v", resp.Title)
	}
}
