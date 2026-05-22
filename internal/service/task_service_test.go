package service

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"tryit.me/internal/app"
	"tryit.me/internal/model"
	"tryit.me/internal/repository"
)

func setup() *TaskService {
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
	return NewTaskService(repo)
}

func IdConversion(id string, t *testing.T) int {
	rId, err := strconv.Atoi(id)
	if err != nil {
		t.Fatalf("expected numeric id got: %v", err)
		return 0
	}
	return rId
}

func TestCreateTask(t *testing.T) {
	taskService := setup()

	// acquire
	got, err := taskService.CreateTask(context.Background(), "task checker")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}
	rID, err := strconv.Atoi(got.ID)
	if err != nil {
		t.Fatalf("expected valid numeric ID: %v", err)
	}

	// release
	t.Cleanup(func() {
		_, err := taskService.DeleteTask(context.Background(), rID)
		if err != nil {
			t.Errorf("cleanup failed deleting task %d: %v", rID, err)
		}
	})

	if got.Title != "task checker" {
		t.Fatalf("expected title 'task checker', got '%s'", got.Title)
	}

}

func TestCreateTask_Without_Title(t *testing.T) {
	taskService := setup()

	_, createErr := taskService.CreateTask(context.Background(), "")
	if createErr == nil {
		t.Fatal("expected error on empty or missing title")
	}
}

func TestGetAll(t *testing.T) {
	taskService := setup()

	ftsk, _ := taskService.CreateTask(context.Background(), "task1")
	stsk, _ := taskService.CreateTask(context.Background(), "task2")

	fID := IdConversion(ftsk.ID, t)
	sId := IdConversion(stsk.ID, t)

	t.Cleanup(func() {
		delete := func(id int) {
			if _, err := taskService.DeleteTask(context.Background(), id); err != nil {
				t.Errorf("cleanup failed deleting task %d: %v", id, err)
			}
		}

		delete(fID)
		delete(sId)
	})

	got, err := taskService.GetAll(context.Background())
	if err != nil {
		t.Fatalf("expected fetched for task to succed: %v", err)
	}

	if len(got) < 2 {
		t.Fatalf("got %d tasks, expected >= 2", len(got))
	}
}

func TestGetTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask(context.Background(), "checkOne")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	id, err := strconv.Atoi(task.ID)
	if err != nil {
		t.Fatalf("expeted valid format ID %v", err)
	}

	t.Cleanup(func() {
		_, err := taskService.DeleteTask(context.Background(), id)
		if err != nil {
			t.Errorf("cleanup failed on deleting task %d: %v", id, err)
		}
	})

	got, err := taskService.GetTask(context.Background(), id)
	if err != nil {
		t.Fatalf("expected task to exist at the least: %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("got ID %s, expected %s", got.ID, task.ID)
	}

	if got.Title != task.Title {
		t.Errorf("got Title %s, expected %s", got.Title, task.Title)
	}
}

func TestGetTask_WithoutID(t *testing.T) {
	taskService := setup()

	_, err := taskService.GetTask(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for missing task")
	}
}

func TestDeleteTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask(context.Background(), "Checkers")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	id, err := strconv.Atoi(task.ID)
	if err != nil {
		t.Fatalf("expected valid ID format: %v", err.Error())
	}

	got, err := taskService.DeleteTask(context.Background(), id)
	if err != nil {
		t.Fatalf("expected task deletion to succeed: %v", err)
	}

	if got.ID != task.ID {
		t.Fatalf("got ID %s, expected ID %s", got.ID, task.ID)
	}
}

func TestDeleteTask_WithoutID(t *testing.T) {
	taskService := setup()

	_, err := taskService.DeleteTask(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for missing task ID")
	}
}

func TestUpdateTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask(context.Background(), "Hella Chao")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	id, err := strconv.Atoi(task.ID)
	if err != nil {
		t.Fatalf("expecte valid format ID: %v", err)
	}

	t.Cleanup(func() {
		_, err := taskService.DeleteTask(context.Background(), id)
		if err != nil {
			t.Errorf("cleanup failed for taskID %d: %v", id, err)
		}
	})

	updated_task := model.UpdateTaskInput{
		IsCompleted: true,
		Title:       "Hella",
	}

	got, err := taskService.UpdateTask(context.Background(), id, updated_task)
	if err != nil {
		t.Fatal("expected task updated to succeed")
	}

	if got.IsCompleted != updated_task.IsCompleted {
		t.Errorf("got IsCompleted: %t, expected IsCompleted: %t", got.IsCompleted, updated_task.IsCompleted)
	}

	if got.Title != updated_task.Title {
		t.Errorf("got titile %s, expected %s", got.Title, updated_task.Title)
	}
}

func TestUpdateTask_Without_ID_Task(t *testing.T) {
	taskService := setup()

	task := model.UpdateTaskInput{
		Title: "check this",
	}

	_, err := taskService.UpdateTask(context.Background(), 0, task)
	if err == nil {
		t.Fatal("expected error on missing id")
	}

	_, taskErr := taskService.UpdateTask(context.Background(), 0, model.UpdateTaskInput{})
	if taskErr == nil {
		t.Fatal("expected error on missing task and id")
	}
}
