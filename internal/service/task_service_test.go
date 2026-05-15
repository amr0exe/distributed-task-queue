package service

import (
	"testing"

	"tryit.me/internal/model"
	"tryit.me/internal/store"
)

func setup() *TaskService {
	st := store.NewMemoryStore()
	return NewTaskService(st)
}

func TestCreateTask(t *testing.T) {
	taskService := setup()

	got, err := taskService.CreateTask("task checker")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	if got.Title != "task checker" {
		t.Fatalf("expected title 'task checker', got '%s'", got.Title)
	}
}

func TestCreateTask_Without_Title(t *testing.T) {
	taskService := setup()

	_, createErr := taskService.CreateTask("")
	if createErr == nil {
		t.Fatal("expected error on empty or missing title")
	}
}

func TestGetAll(t *testing.T) {
	taskService := setup()

	_, _ = taskService.CreateTask("task1")
	_, _ = taskService.CreateTask("task2")

	got, err := taskService.GetAll()
	if err != nil {
		t.Fatalf("expected to be fetched: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("got %d tasks, expected 2", len(got))
	}
}

func TestGetTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask("checkOne")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	got, err := taskService.GetTask(task.ID)
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

	_, err := taskService.GetTask("hahah")
	if err == nil {
		t.Fatal("expected error for missing task")
	}
}

func TestDeleteTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask("Checkers")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	got, err := taskService.DeleteTask(task.ID)
	if err != nil {
		t.Fatalf("expected task deletion to succeed: %v", err)
	}

	if got.ID != task.ID {
		t.Fatalf("got ID %s, expected ID %s", got.ID, task.ID)
	}
}

func TestDeleteTask_WithoutID(t *testing.T) {
	taskService := setup()

	_, err := taskService.DeleteTask("wrong_it_is")
	if err == nil {
		t.Fatal("expected error for missing task ID")
	}
}

func TestUpdateTask(t *testing.T) {
	taskService := setup()

	task, err := taskService.CreateTask("Hella Chao")
	if err != nil {
		t.Fatalf("expected task creation to succeed: %v", err)
	}

	updated_task := model.Task{
		ID:          "33",
		IsCompleted: true,
		Title:       "Hella",
	}

	got, err := taskService.UpdateTask(task.ID, updated_task)
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

	task := model.Task{
		Title: "check this",
	}

	_, err := taskService.UpdateTask("", task)
	if err == nil {
		t.Fatal("expected error on missing id")
	}

	_, taskErr := taskService.UpdateTask("", model.Task{})
	if taskErr == nil {
		t.Fatal("expected error on missing task and id")
	}
}
