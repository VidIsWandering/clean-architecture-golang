// Package dto contains Data Transfer Objects for the application layer.
// These DTOs are used to transfer data between use cases and external layers.
package dto

// CreateTaskRequest represents the input data for creating a new task.
type CreateTaskRequest struct {
	Title       string
	Description string
}
