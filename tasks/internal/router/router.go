package router

import (
	"net/http"
)

type Handlers interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	PostTask(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	r Handlers
}

func NewRouter(a Handlers) *Router {
	return &Router{
		r: a,
	}
}

func (router *Router) InitRoutes(MultiPlex *http.ServeMux) {

	MultiPlex.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			router.r.GetTask(w, r)
		case http.MethodDelete:
			router.r.DeleteTask(w, r)
		default:
			http.Error(w, "Данный метод не разрешён или не существует", http.StatusMethodNotAllowed)
		}

	})
	MultiPlex.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			router.r.PostTask(w, r)
		default:
			http.Error(w, "Данный метод не разрешён или не существует", http.StatusMethodNotAllowed)
		}
	})
}
