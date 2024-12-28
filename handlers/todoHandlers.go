package handlers

import (
	"ToDoListAPI/models"
	"ToDoListAPI/utils"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// CreateNewListHandler TODO
func CreateNewListHandler(db *sql.DB) http.HandlerFunc {
	type Request struct {
		todoListName string `json:"name"`
	}

	type Response struct {
		todoListName string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get name from url parameter
		newToDoListName := chi.URLParam(r, "todoListName")

		// TODO store new list in db
		if err := models.InsertNewTodo(db, newToDoListName); err != nil {
			fmt.Println(err)
			utils.JSONResponse(w, http.StatusInternalServerError, nil)
		}

		// TODO return created
		utils.JSONResponse(w, http.StatusCreated, nil)
	}
}

// AddToListHandler TODO
func AddToListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetListHandler TODO
func GetListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// RemoveFromListHandler TODO
func RemoveFromListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
