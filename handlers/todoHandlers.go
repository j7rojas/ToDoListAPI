package handlers

import (
	"ToDoListAPI/models"
	"ToDoListAPI/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// CreateNewListHandler TODO
func CreateNewListHandler(db *sql.DB) http.HandlerFunc {
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
	// TODO Define Request Struct
	// Todo Define Repsonse Struct

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get name of the TodoList from url parameter
		newToDoListName := chi.URLParam(r, "todoListName")

		var newTask models.Task
		// TODO get the task information to add from json
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// TODO Look up the TodoList id from db
		models.GetTodoList(db, newToDoListName)

		// TODO Insert new task to db
		//if err = models.InsertNewTask(db); err != nil {
		//	fmt.Println(err)
		//	utils.JSONResponse(w, http.StatusInternalServerError, nil)
		//}

		// TODO return updated TodoList
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
