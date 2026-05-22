package service

// services will be responsible for interacting with DB/IN-MeM

import (
	"context"
	"errors"
	"fmt"

	"tryit.me/internal/model"
	"tryit.me/internal/repository"
)

type TaskService struct {
	db *repository.TaskRepository
}

func NewTaskService(db *repository.TaskRepository) *TaskService {
	return &TaskService{
		db: db,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, title string) (model.Task, error) {
	if title == "" {
		return model.Task{}, errors.New("title cannot be empty")
	}

	rTsk, err := s.db.CreateTask(ctx, title)
	if err != nil {
		return model.Task{}, fmt.Errorf("CreateTask failed: %w", err)
	}
	fmt.Printf("Task Creation successfull. %s \n", rTsk.Title)

	return rTsk, nil
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.db.ListTasks(ctx)
	if err != nil {
		return []model.Task{}, fmt.Errorf("GetAll failed: %w", err)
	}
	return tasks, nil
}

func (s *TaskService) GetTask(ctx context.Context, id int) (model.Task, error) {
	if id == 0 {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, err := s.db.GetTask(ctx, id)
	if err != nil {
		return model.Task{}, fmt.Errorf("Failed fetching for taskID: %d", id)
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) (model.Task, error) {
	if id == 0 {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, err := s.db.DeleteTask(ctx, id)
	if err != nil {
		return model.Task{}, fmt.Errorf("Failed deleting task with taskID: %d", id)
	}

	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, updatedTask model.UpdateTaskInput) (model.Task, error) {
	if id == 0 {
		return model.Task{}, errors.New("ID cannot be empty")
	}

	task, err := s.db.UpdateTask(ctx, id, updatedTask)
	if err != nil {
		return model.Task{}, fmt.Errorf("Failed updating task with taskID: %d", id)
	}

	return task, nil
}
