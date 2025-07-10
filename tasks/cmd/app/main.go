package main

import (
	"log"
	"net/http"
	"tasks/internal/handlers"
	"tasks/internal/router"
	"tasks/internal/service"
	"tasks/internal/storage"
)

func main() {

	store := storage.NewStorage()

	svc := service.NewService(store)

	h := handlers.NewHandler(svc)

	mux := http.NewServeMux()

	router := router.NewRouter(h)

	router.InitRoutes(mux)
	log.Println("Слушаю на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
