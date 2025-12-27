package entities

import (
	"clean-architecture-golang/domain/value_objects"
	"testing"
)

func TestNewTask(t *testing.T) {
	title := "Test Task"
	description := "Test Description"

	task, err := NewTask(title, description)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if task.Title != title {
		t.Errorf("Expected title %s, got %s", title, task.Title)
	}

	if task.Description != description {
		t.Errorf("Expected description %s, got %s", description, task.Description)
	}

	if task.Status.String() != "todo" {
		t.Errorf("Expected status todo, got %s", task.Status.String())
	}

	if task.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestNewTask_EmptyTitle(t *testing.T) {
	_, err := NewTask("", "description")
	if err == nil {
		t.Error("Expected error for empty title")
	}
}

func TestUpdateStatus(t *testing.T) {
	task, _ := NewTask("Test", "Desc")

	// Valid transition
	err := task.UpdateStatus(value_objects.TaskStatus("doing"))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if task.Status.String() != "doing" {
		t.Errorf("Expected status doing, got %s", task.Status.String())
	}

	// Invalid status
	err = task.UpdateStatus(value_objects.TaskStatus("invalid"))
	if err == nil {
		t.Error("Expected error for invalid status")
	}

	// Invalid transition: done -> todo
	task.Status = value_objects.TaskStatus("done")
	err = task.UpdateStatus(value_objects.TaskStatus("todo"))
	if err == nil {
		t.Error("Expected error for done to todo transition")
	}
}
