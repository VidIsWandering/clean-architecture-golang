package repositories

import (
	"errors"
	"testing"

	"clean-architecture-golang/domain/entities"
	"clean-architecture-golang/domain/value_objects"
)

func TestSaveFindDeleteFlow(t *testing.T) {
	r := NewInMemoryTaskRepository()

	// create domain task
	task, err := entities.NewTask("title", "desc")
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	// Save
	if err := r.Save(task); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	// FindById
	f, err := r.FindById(task.ID)
	if err != nil {
		t.Fatalf("find by id failed: %v", err)
	}
	if f.Title != task.Title || f.Description != task.Description {
		t.Fatalf("found task mismatch")
	}

	// FindByStatus
	list, err := r.FindByStatus(value_objects.StatusTodo)
	if err != nil {
		t.Fatalf("find by status failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 task, got %d", len(list))
	}

	// Delete
	if err := r.Delete(task.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	// FindById afterwards -> ErrNotFound
	_, err = r.FindById(task.ID)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}
