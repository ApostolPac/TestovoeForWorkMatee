package storage

import (
	"sync"
	"tasks/internal/models"
)

type Storage struct {
	sync.RWMutex
	tasks map[string]*models.Tasks
}

func NewStorage()*Storage{
	return &Storage{
		tasks: make(map[string]*models.Tasks),
	}
}

func (s *Storage) GetTask(uuid string) (*models.Tasks, bool) {
	s.RLock()
	task, exists := s.tasks[uuid]
	defer s.RUnlock()
	return task, exists
}

func (s *Storage) PostTask(task *models.Tasks){
	s.Lock()
	s.tasks[task.Uuid] = task
	defer s.Unlock()
}

func (s *Storage) DeleteTask(uuid string){
	s.Lock()
	delete(s.tasks, uuid)
	defer s.Unlock()
}
