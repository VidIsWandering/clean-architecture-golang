package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"clean-architecture-golang/domain/value_objects"
	testutil "clean-architecture-golang/internal/testutil"
)

func TestCreate_EmptyTitle_Returns400JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	reqBody := map[string]string{"title": "", "description": "desc"}
	body, _ := json.Marshal(reqBody)
	resp, err := http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}

func TestUpdateStatus_InvalidStatus_Returns400JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	createResp := testutil.CreateTask(t, server.URL, "ok", "desc")
	taskID := createResp["ID"].(string)

	updateBody := map[string]string{"newStatus": "invalid"}
	updateJSON, _ := json.Marshal(updateBody)
	req, _ := http.NewRequest("PUT", server.URL+"/tasks/"+taskID+"/status", bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	updateResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer updateResp.Body.Close()
	if updateResp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", updateResp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(updateResp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}

func TestDelete_NotFound_Returns404JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	validId := string(value_objects.NewTaskId())
	req, _ := http.NewRequest("DELETE", server.URL+"/tasks/"+validId, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}

func TestUpdateStatus_EmptyID_Returns400JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	updateBody := map[string]string{"newStatus": "doing"}
	updateJSON, _ := json.Marshal(updateBody)
	// send request with an empty segment
	req, _ := http.NewRequest("PUT", server.URL+"/tasks/%20/status", bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	updateResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer updateResp.Body.Close()
	if updateResp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", updateResp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(updateResp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}

func TestGetTasksByStatus_InvalidStatus_Returns400JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks?status=invalid")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}

func TestUpdateStatus_NotFound_Returns404JSON(t *testing.T) {
	server, _ := testutil.SetupTestServer()
	defer server.Close()

	updateBody := map[string]string{"newStatus": "doing"}
	updateJSON, _ := json.Marshal(updateBody)
	validId := string(value_objects.NewTaskId())
	req, _ := http.NewRequest("PUT", server.URL+"/tasks/"+validId+"/status", bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	updateResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer updateResp.Body.Close()
	if updateResp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", updateResp.StatusCode)
	}
	var errResp map[string]string
	if err := json.NewDecoder(updateResp.Body).Decode(&errResp); err != nil {
		t.Fatalf("failed decode: %v", err)
	}
	if _, ok := errResp["error"]; !ok {
		t.Fatalf("expected error field in response")
	}
}
