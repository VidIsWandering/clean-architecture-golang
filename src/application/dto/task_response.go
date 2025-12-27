package dto

import (
	"clean-architecture-golang/domain/entities"
	"time"
)

// TaskResponse represents the output data for task operations.
// It prevents leaking domain entities to the presentation layer.
type TaskResponse struct {
	ID          string
	Title       string
	Status      string
	Description string
	CreatedAt   string
}

// ToTaskResponse converts a domain Task entity to a TaskResponse DTO.
func ToTaskResponse(t *entities.Task) TaskResponse {
	return TaskResponse{
		ID:          string(t.ID),
		Title:       t.Title,
		Status:      t.Status.String(),
		Description: t.Description,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
	}
}
