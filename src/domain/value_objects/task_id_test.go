package value_objects

import "testing"

func TestNewTaskId(t *testing.T) {
	id1 := NewTaskId()
	id2 := NewTaskId()

	if id1 == "" {
		t.Error("Expected non-empty ID")
	}

	if id2 == "" {
		t.Error("Expected non-empty ID")
	}

	if id1 == id2 {
		t.Error("Expected unique IDs")
	}

	// Validate ID format using ParseTaskId (UUID)
	if _, err := ParseTaskId(string(id1)); err != nil {
		t.Fatalf("ParseTaskId failed: %v", err)
	}
}
