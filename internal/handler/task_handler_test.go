package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"tryit.me/internal/service"
	"tryit.me/internal/store"
)

func setup() *service.TaskService {
	st := store.NewMemoryStore()
	return service.NewTaskService(st)
}

func TestCreateTask(t *testing.T) {
	tskSrv := setup()
	handler := NewTaskHandler(tskSrv)

	body := `{ "title": "learn go" }`

	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.CreateTask(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Fatalf("got status %d, expected %d", rr.Code, http.StatusAccepted)
	}

	var got map[string]any

	err := json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed decoding response body: %v", err)
	}

	if got["title"] != "learn go" {
		t.Fatalf("got title %v, expected %s", got["title"], "learn go")
	}
}

func TestCreateTask_EmptyBody(t *testing.T) {
	tskSrv := setup()
	handler := NewTaskHandler(tskSrv)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.CreateTask(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, expected %d", rr.Code, http.StatusBadRequest)
	}
}
