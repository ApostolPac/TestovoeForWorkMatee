package service

import (
	"fmt"
	"tasks/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

type Storage interface {
	PostTask(task *models.Tasks)
	DeleteTask(uuid string)
	GetTask(uuid string) (*models.Tasks, bool)
}

type Service struct {
	storage Storage
}

func NewService(a Storage) *Service {
	return &Service{
		storage: a,
	}
}

func generateUuid() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (s *Service) GetTask(uuid string) (*models.Tasks, error) {
	task, ok := s.storage.GetTask(uuid)
	if !ok {
		return nil, fmt.Errorf("задача не была найдена")
	}
	return task, nil
}

func (s *Service) DeleteTask(uuid string) error {
	_, ok := s.storage.GetTask(uuid)
	if !ok {
		return fmt.Errorf("задача не была найдена")
	}
	s.storage.DeleteTask(uuid)
	return nil
}

func (s *Service) PostTask(task models.Tasks) (*models.Tasks, error) {
	uuid, err := generateUuid()
	if err != nil {
		return nil, err
	}
	task.Uuid = uuid
	task.CreatedAt = time.Now()
	task.Status = models.StatusRunning
	rand.Seed(uint64(time.Now().UnixNano()))

	s.storage.PostTask(&task)

	go func(t *models.Tasks) {
		start := time.Now()

		t.StartedAt = &start
		delay := time.Duration(3+rand.Intn(3)) * time.Minute
		time.Sleep(delay)

		finish := time.Now()

		t.FinishedAt = &finish

		t.TimeForProcessing = fmt.Sprintf("%v min", int(finish.Sub(start).Minutes()))
		t.Status = models.StatusCompleted

		s.storage.PostTask(t)
	}(&task)

	return &task, nil
}
