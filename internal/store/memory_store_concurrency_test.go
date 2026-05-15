package store

import (
	"strconv"
	"sync"
	"testing"

	"tryit.me/internal/model"
)

func TestMemoryStore_ConcurrentCreate(t *testing.T) {
	st := NewMemoryStore()

	var wg sync.WaitGroup
	count := 1000

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			task := model.Task{
				ID:    strconv.Itoa(i),
				Title: "task",
			}

			st.Create(task)
		}(i)
	}

	wg.Wait()

	tasks, _ := st.List()

	if len(tasks) != count {
		t.Fatalf("expected %d tasks, got %d", count, len(tasks))
	}
}

func TestMemoryStore_ConcurrentReadWrite(t *testing.T) {
	st := NewMemoryStore()

	var wg sync.WaitGroup
	count := 1000

	for i := 0; i < count; i++ {
		wg.Add(2)

		go func(i int) {
			defer wg.Done()

			task := model.Task{
				ID:    strconv.Itoa(i),
				Title: "task",
			}

			st.Create(task)
		}(i)

		go func(i int) {
			defer wg.Done()

			st.Get(strconv.Itoa(i))
		}(i)
	}

	wg.Wait()
}
