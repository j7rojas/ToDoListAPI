package routes

import (
	"ToDoListAPI/handlers"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func SetUpRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	SetUpRoutes(r, db)
	return r
}

func SetUpRoutes(r *chi.Mux, db *sql.DB) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/{todoListName}", handlers.CreateNewListHandler(db))
	r.Post("/add/{todoListName}", handlers.AddToListHandler(db))
	r.Get("/{todoListName}", handlers.GetListHandler(db))
	r.Delete("/{todoListName}", handlers.RemoveFromListHandler(db))
}
