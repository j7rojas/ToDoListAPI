package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// CreateNewListHandler TODO
func CreateNewListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get name from url parameter
		newToDoListName := chi.URLParam(r, "todoListName")

		// TODO store new list in db
		InsertTodo(db, newToDoListName, false)

		// TODO return created
		w.WriteHeader(http.StatusCreated)
	}
}

// AddToListHandler TODO
func AddToListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetListHandler TODO
func GetListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// RemoveFromListHandler TODO
func RemoveFromListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
