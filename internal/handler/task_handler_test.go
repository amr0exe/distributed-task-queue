package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"tryit.me/internal/app"
	"tryit.me/internal/repository"
	"tryit.me/internal/service"
)

func setup() *service.TaskService {
	_ = godotenv.Load("../../.env")
	dbStr, ok := os.LookupEnv("TEST_DBSTRING")
	if !ok {
		log.Fatalf("testing db string not found")
	}

	a, err := app.New(app.Config{DBString: dbStr})
	if err != nil {
		log.Fatalf("failed initializing app: %v", err.Error())
	}

	repo := repository.NewTaskRepository(a.DB())
	return service.NewTaskService(repo)
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
