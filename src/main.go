// Package main is the entry point of the application.
// It wires all dependencies using dependency injection and starts the HTTP server.
package main

import (
	"clean-architecture-golang/application/usecases"
	"clean-architecture-golang/infrastructure/repositories"
	"clean-architecture-golang/presentation/controllers"
	"net/http"
	"strings"
)

func main() {
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
		if strings.HasSuffix(r.URL.Path, "/status") && r.Method == http.MethodPut {
			controller.UpdateStatus(w, r)
		} else if r.Method == http.MethodDelete {
			controller.Delete(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", mux)
}
