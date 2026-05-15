package service

// services will be responsible for interacting with DB/IN-MeM

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"tryit.me/internal/model"
	"tryit.me/internal/store"
)

type TaskService struct {
	// in-memory instance to share between services
	store *store.MemoryStore
}

func NewTaskService(store *store.MemoryStore) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) CreateTask(title string) (model.Task, error) {
	if title == "" {
		return model.Task{}, errors.New("title cannot be empty")
	}

	task := model.Task{
		ID:          uuid.NewString(),
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	s.store.Create(task)
	fmt.Printf("Task Creation successfull. %s \n", task.Title)

	return task, nil
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	tasks, _ := s.store.List()
	return tasks, nil
}

func (s *TaskService) GetTask(id string) (model.Task, error) {
	if id == "" {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, ok := s.store.Get(id)
	if !ok {
		return model.Task{}, fmt.Errorf("Failed fetching for taskID: %s", id)
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) (model.Task, error) {
	if id == "" {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, ok := s.store.Delete(id)
	if !ok {
		return model.Task{}, fmt.Errorf("Failed deleting task with taskID: %s", id)
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask model.Task) (model.Task, error) {
	if id == "" {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, ok := s.store.Update(id, updatedTask)
	if !ok {
		return model.Task{}, fmt.Errorf("Failed updating task with taskID: %s", id)
	}

	return task, nil
}
