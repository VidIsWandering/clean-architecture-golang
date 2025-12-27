package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"clean-architecture-golang/application/usecases"
	"clean-architecture-golang/infrastructure/repositories"
	"clean-architecture-golang/presentation/controllers"
)

func setupTestServer() *httptest.Server {
	repo := repositories.NewInMemoryTaskRepository()

	createUC := &usecases.CreateTaskUseCase{Repo: repo}
	updateUC := &usecases.UpdateTaskStatusUseCase{Repo: repo}
	getUC := &usecases.GetTasksByStatusUseCase{Repo: repo}
	deleteUC := &usecases.DeleteTaskUseCase{Repo: repo}

	controller := &controllers.TaskController{
		CreateTaskUC:       createUC,
		UpdateStatusUC:     updateUC,
		GetTasksByStatusUC: getUC,
		DeleteTaskUC:       deleteUC,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.Create(w, r)
		case http.MethodGet:
			controller.ListByStatus(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if bytes.HasSuffix([]byte(r.URL.Path), []byte("/status")) && r.Method == http.MethodPut {
			controller.UpdateStatus(w, r)
		} else if r.Method == http.MethodDelete {
			controller.Delete(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return httptest.NewServer(mux)
}

func TestCreateTask(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	reqBody := map[string]string{
		"title":       "Test Task",
		"description": "Test Description",
	}
	body, _ := json.Marshal(reqBody)

	resp, err := http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["Title"] != "Test Task" {
		t.Errorf("Expected title 'Test Task', got %v", response["Title"])
	}
}

func TestGetTasksByStatus(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// First create a task
	reqBody := map[string]string{
		"title":       "Test Task",
		"description": "Test Description",
	}
	body, _ := json.Marshal(reqBody)
	http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(body))

	// Then get tasks by status
	resp, err := http.Get(server.URL + "/tasks?status=todo")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(response) != 1 {
		t.Errorf("Expected 1 task, got %d", len(response))
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create a task first
	reqBody := map[string]string{
		"title":       "Test Task",
		"description": "Test Description",
	}
	body, _ := json.Marshal(reqBody)
	resp, _ := http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(body))

	var createResp map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&createResp)
	taskID := createResp["ID"].(string)
	resp.Body.Close()

	// Update status
	updateBody := map[string]string{"newStatus": "doing"}
	updateJSON, _ := json.Marshal(updateBody)
	req, _ := http.NewRequest("PUT", server.URL+"/tasks/"+taskID+"/status", bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	updateResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer updateResp.Body.Close()

	if updateResp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", updateResp.StatusCode)
	}
}

func TestDeleteTask(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create a task first
	reqBody := map[string]string{
		"title":       "Test Task",
		"description": "Test Description",
	}
	body, _ := json.Marshal(reqBody)
	resp, _ := http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(body))

	var createResp map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&createResp)
	taskID := createResp["ID"].(string)
	resp.Body.Close()

	// Delete task
	req, _ := http.NewRequest("DELETE", server.URL+"/tasks/"+taskID, nil)
	client := &http.Client{}
	deleteResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer deleteResp.Body.Close()

	if deleteResp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", deleteResp.StatusCode)
	}
}
