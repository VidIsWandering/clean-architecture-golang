package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"clean-architecture-golang/application/usecases"
	"clean-architecture-golang/infrastructure/repositories"
	presentation "clean-architecture-golang/presentation/controllers"
)

// SetupTestServer creates an httptest.Server wired with InMemoryTaskRepository
// and returns the server and the repository for further inspection.
func SetupTestServer() (*httptest.Server, *repositories.InMemoryTaskRepository) {
	repo := repositories.NewInMemoryTaskRepository()

	createUC := &usecases.CreateTaskUseCase{Repo: repo}
	updateUC := &usecases.UpdateTaskStatusUseCase{Repo: repo}
	getUC := &usecases.GetTasksByStatusUseCase{Repo: repo}
	deleteUC := &usecases.DeleteTaskUseCase{Repo: repo}

	controller := &presentation.TaskController{
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
		if r.Method == http.MethodPut && len(r.URL.Path) > 0 && len(r.URL.Path) >= 7 && r.URL.Path[len(r.URL.Path)-7:] == "/status" {
			controller.UpdateStatus(w, r)
		} else if r.Method == http.MethodDelete {
			controller.Delete(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return httptest.NewServer(mux), repo
}

// CreateTask helper posts to /tasks and returns the created task response as map.
func CreateTask(t *testing.T, serverURL string, title, description string) map[string]interface{} {
	t.Helper()
	reqBody := map[string]string{"title": title, "description": description}
	body, _ := json.Marshal(reqBody)
	resp, err := http.Post(serverURL+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("decode create response failed: %v", err)
	}
	return result
}
