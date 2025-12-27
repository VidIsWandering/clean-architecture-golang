package value_objects

// TaskStatus represents the possible states of a task.
type TaskStatus string

const (
	StatusTodo  TaskStatus = "todo"
	StatusDoing TaskStatus = "doing"
	StatusDone  TaskStatus = "done"
)

// String returns the string representation of the status.
func (s TaskStatus) String() string {
	return string(s)
}

// IsValid checks if the status is one of the allowed values.
func (s TaskStatus) IsValid() bool {
	switch s {
	case StatusTodo, StatusDoing, StatusDone:
		return true
	default:
		return false
	}
}
