package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"tryit.me/internal/model"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		pool: pool,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, title string) (model.Task, error) {
	var t model.Task

	query := `INSERT INTO tasks (title) VALUES ($1) RETURNING id, title, is_completed, created_at`
	err := r.pool.QueryRow(ctx, query, title).Scan(&t.ID, &t.Title, &t.IsCompleted, &t.CreatedAt)

	if err != nil {
		return model.Task{}, fmt.Errorf("repo.CreateTask failed: %w", err)
	}

	return t, nil
}

func (r *TaskRepository) ListTasks(ctx context.Context) ([]model.Task, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, is_completed, created_at
		FROM tasks
		ORDER BY id ASC`)
	if err != nil {
		return []model.Task{}, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.IsCompleted, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("repo.ListTasks scan failed: %w", err)
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.ListTasks rows failed: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetTask(ctx context.Context, id int) (model.Task, error) {
	var t model.Task
	err := r.pool.QueryRow(ctx,
		`SELECT id, title, is_completed, created_at
		FROM tasks
		WHERE id = $1`,
		id,
	).Scan(&t.ID, &t.Title, &t.IsCompleted, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Task{}, fmt.Errorf("repo.GetTask failed: task %d not found", id)
	}
	if err != nil {
		return model.Task{}, fmt.Errorf("repo.GetTask: %w", err)
	}

	return t, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id int, in model.UpdateTaskInput) (model.Task, error) {
	var t model.Task
	err := r.pool.QueryRow(ctx,
		`UPDATE tasks
		SET
			title = COALESCE($1, title),
			is_completed = COALESCE($2, is_completed)
		WHERE id = $3
		RETURNING id, title, is_completed, created_at`,
		in.Title, in.IsCompleted,
		id,
	).Scan(&t.ID, &t.Title, &t.IsCompleted, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows){
		return model.Task{}, fmt.Errorf("repo.UpdateTask: task %d not found", id)
	}
	if err != nil {
		return model.Task{}, fmt.Errorf("repo.UpdateTask: %w", err)
	}

	return t, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int) (model.Task, error) {
	var t model.Task
	err := r.pool.QueryRow(ctx,
		`DELETE FROM tasks
		WHERE id = $1
		RETURNING id, title, is_completed, created_at `,
		id,
	).Scan(&t.ID, &t.Title, &t.IsCompleted, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Task{}, fmt.Errorf("repo.DeleteTask: task %d not found", id)
	}
	if err != nil {
		return model.Task{}, fmt.Errorf("repo.DeleteTask: %w", err)
	}

	return t, nil
}
