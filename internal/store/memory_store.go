package store

import (
	"sync"

	"tryit.me/internal/model"
)

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]model.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]model.Task),
	}
}

func (s *MemoryStore) Create(task model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[task.ID] = task
}

func (s *MemoryStore) List() ([]model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]model.Task, 0, len(s.data))

	for _, task := range s.data {
		tasks = append(tasks, task)
	}

	return tasks, true
}

func (s *MemoryStore) Get(id string) (model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.data[id]

	return task, ok
}

func (s *MemoryStore) Delete(id string) (model.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.data[id]
	if !ok {
		return model.Task{}, false
	}

	delete(s.data, id)

	return task, true
}

func (s *MemoryStore) Update(id string, updatedTask model.Task) (model.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[id]
	if !ok {
		return model.Task{}, false
	}

	updatedTask.ID = id

	s.data[id] = updatedTask

	return updatedTask, true
}
