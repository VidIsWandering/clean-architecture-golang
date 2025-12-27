// Package value_objects contains value objects that represent immutable domain concepts.
// These objects have their own validation and behavior.
package value_objects

import (
	"errors"

	"github.com/google/uuid"
)

// TaskId represents a unique identifier for a task.
// It uses UUID (RFC 4122) string format.
type TaskId string

// NewTaskId generates a new UUID-based TaskId.
func NewTaskId() TaskId {
	return TaskId(uuid.NewString())
}

// ErrInvalidTaskId indicates the provided task id is invalid or malformed.
var ErrInvalidTaskId = errors.New("invalid task id")

// ParseTaskId validates and parses a string into a TaskId.
func ParseTaskId(s string) (TaskId, error) {
	if s == "" {
		return "", ErrInvalidTaskId
	}
	if _, err := uuid.Parse(s); err != nil {
		return "", ErrInvalidTaskId
	}
	return TaskId(s), nil
}
