package models

import "time"

type Status string

const (
	StatusRunning   Status = "обрабатывается"
	StatusCompleted Status = "завершена"
	StatusFailed    Status = "ошибка при обработке"
)

type Tasks struct {
	Uuid              string     `json:"uuid"`
	Status            Status     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	StartedAt         *time.Time `json:"started_at,omitempty"`
	FinishedAt        *time.Time `json:"finished_at,omitempty"`
	TimeForProcessing string `json:"time_for_processing"`
}
