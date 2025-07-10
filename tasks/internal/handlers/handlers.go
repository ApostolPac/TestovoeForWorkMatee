package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"tasks/internal/models"
)

type Service interface {
	GetTask(uuid string) (*models.Tasks, error)
	DeleteTask(uuid string) error
	PostTask(task models.Tasks) (*models.Tasks, error)
}

type Handlers struct {
	s Service
}

func NewHandler(a Service) *Handlers {
	return &Handlers{
		s: a,
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {

	uuid := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if uuid == "" {
		http.NotFound(w, r)
		return
	}
	task, err := h.s.GetTask(uuid)
	if err != nil {
		http.Error(w, "Задача не была найдена", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if uuid == "" {
		http.NotFound(w, r)
		return
	}
	err := h.s.DeleteTask(uuid)
	if err != nil {
		http.Error(w, "Задача не была удалена", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, "Задача успешно удалена")
}

func (h *Handlers) PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Неподходящий формат", http.StatusBadRequest)
		return
	}
	createdTask, err := h.s.PostTask(task)
	if err != nil {
		http.Error(w, "Ошибка при создании задачи", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusCreated, createdTask)
}
