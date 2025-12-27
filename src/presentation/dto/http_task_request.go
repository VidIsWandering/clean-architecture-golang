// Package dto contains Data Transfer Objects specific to the presentation layer.
// These DTOs handle HTTP request/response formats and can include framework-specific tags.
package dto

// HttpCreateTaskRequest represents the JSON payload for creating a task via HTTP.
type HttpCreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// HttpUpdateStatusRequest represents the JSON payload for updating task status via HTTP.
type HttpUpdateStatusRequest struct {
	NewStatus string `json:"newStatus"`
}
