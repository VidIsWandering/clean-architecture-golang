// Package controllers contains HTTP handlers for the presentation layer.
// Controllers handle HTTP requests, perform basic validation, and delegate to use cases.
package controllers

import (
	"clean-architecture-golang/application/dto"
	"clean-architecture-golang/application/usecases"
	domain_entities "clean-architecture-golang/domain/entities"
	repo "clean-architecture-golang/infrastructure/repositories"
	presentation_dto "clean-architecture-golang/presentation/dto"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

type TaskController struct {
	CreateTaskUC       *usecases.CreateTaskUseCase
	UpdateStatusUC     *usecases.UpdateTaskStatusUseCase
	GetTasksByStatusUC *usecases.GetTasksByStatusUseCase
	DeleteTaskUC       *usecases.DeleteTaskUseCase
}

func writeJSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func (c *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	var httpReq presentation_dto.HttpCreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&httpReq); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	appReq := dto.CreateTaskRequest{
		Title:       httpReq.Title,
		Description: httpReq.Description,
	}
	response, err := c.CreateTaskUC.Execute(appReq)
	if err != nil {
		switch {
		case errors.Is(err, domain_entities.ErrEmptyTitle), errors.Is(err, domain_entities.ErrInvalidStatus):
			writeJSONError(w, http.StatusBadRequest, err.Error())
		default:
			log.Printf("Create internal error: %v", err)
			writeJSONError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *TaskController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id = strings.TrimSuffix(id, "/status")
	if id == "" {
		writeJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var httpReq presentation_dto.HttpUpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&httpReq); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := c.UpdateStatusUC.Execute(id, httpReq.NewStatus)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrInvalidID):
			writeJSONError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, domain_entities.ErrInvalidStatus), errors.Is(err, domain_entities.ErrInvalidTransition):
			writeJSONError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, repo.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			log.Printf("UpdateStatus internal error: %v", err)
			writeJSONError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *TaskController) ListByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		writeJSONError(w, http.StatusBadRequest, "status query param required")
		return
	}
	responses, err := c.GetTasksByStatusUC.Execute(status)
	if err != nil {
		switch {
		case errors.Is(err, domain_entities.ErrInvalidStatus):
			writeJSONError(w, http.StatusBadRequest, err.Error())
		default:
			log.Printf("ListByStatus internal error: %v", err)
			writeJSONError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func (c *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		writeJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	err := c.DeleteTaskUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			log.Printf("Delete internal error: %v", err)
			writeJSONError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
