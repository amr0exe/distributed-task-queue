package store

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"tryit.me/internal/model"
)

func TestCreateTask(t *testing.T) {
	store := NewMemoryStore()

	task := model.Task{
		ID:          uuid.NewString(),
		Title:       "Lern",
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	store.Create(task)

	got, exits := store.data[task.ID]

	if !exits {
		t.Fatalf("task was not stored")
	}

	if got.ID != task.ID {
		t.Errorf("got ID %s want ID %s", got.ID, task.ID)
	}

}
