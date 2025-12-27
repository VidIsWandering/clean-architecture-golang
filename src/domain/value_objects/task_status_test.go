package value_objects

import "testing"

func TestTaskStatus_IsValid(t *testing.T) {
	validStatuses := []TaskStatus{StatusTodo, StatusDoing, StatusDone}

	for _, status := range validStatuses {
		if !status.IsValid() {
			t.Errorf("Expected %s to be valid", status)
		}
	}

	invalidStatus := TaskStatus("invalid")
	if invalidStatus.IsValid() {
		t.Error("Expected invalid status to be invalid")
	}
}

func TestTaskStatus_String(t *testing.T) {
	if StatusTodo.String() != "todo" {
		t.Errorf("Expected 'todo', got %s", StatusTodo.String())
	}

	if StatusDoing.String() != "doing" {
		t.Errorf("Expected 'doing', got %s", StatusDoing.String())
	}

	if StatusDone.String() != "done" {
		t.Errorf("Expected 'done', got %s", StatusDone.String())
	}
}
