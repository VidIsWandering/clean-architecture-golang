package value_objects

import (
	"testing"
)

func TestParseNewTaskId(t *testing.T) {
	id := NewTaskId()
	_, err := ParseTaskId(string(id))
	if err != nil {
		t.Fatalf("ParseTaskId failed for NewTaskId: %v", err)
	}
}
