package main

import (
	"crud_it_krasava/config"
	"crud_it_krasava/storage"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"log"
	"net/http"
)

var logger *log.Logger = log.New(os.Stdout, "[USER_API] ", log.LstdFlags)
var handler *Handler

func main() {
	config.SetupDB()

	userStorage := storage.NewPostrgresUser(config.DB)
	handler = NewHandler(userStorage)

	router := setupRouter()

	logger.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/users", func(router chi.Router) {
		router.With(validateInput).Post("/", handler.createUser)
	})
	router.Route("/user/{id}", func(router chi.Router) {
		router.Get("/", handler.getUser)
		router.With(validateInput).Patch("/", handler.updateUser)
	})
	return router
}
