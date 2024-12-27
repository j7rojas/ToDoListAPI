package models

import (
	"database/sql"
	"fmt"
)

// TodoList represents a single TodoList
//
// Fields:
//   - ID: A unique identifier for the Todolist.
//   - NAME: The name of the TodoList
//   - Tasks: The items that will need to be completed to complete a Todolist
//   - COMPLETE: Denotes if the Todolist items are all complete
type TodoList struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Tasks    []Task `json:"items"`
	COMPLETE bool   `json:"complete"`
}

// Task represents a single item that needs to be completed in a TodoList
//
// Fields:
//   - ID: A unique identifier for the Task.
//   - NAME: The name item that needs to be completed
//   - COMPLETE: Denotes if the Task is complete
type Task struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	COMPLETE bool   `json:"complete"`
}

func InsertNewTodo(db *sql.DB, newTodoListName string) error {
	_, err := db.Exec("INSERT INTO todos (name) VALUES (?)", newTodoListName)
	if err != nil {
		return fmt.Errorf("failed to insert todos: %w", err)
	}
	return nil
}

func GetTodoList(db *sql.DB, todoListName string) error {
	result, err := db.Exec("SELECT id, name, complete FROM todos WHERE LOWER(?) = LOWER(name)", todoListName)
	if err != nil {
		return fmt.Errorf("failed to insert todos: %w", err)
	}

	fmt.Println(result)
	return nil
}

func InsertNewTask(db *sql.DB, newTaskName string, todoListId int) error {
	_, err := db.Exec("INSERT INTO todo_items (name, todo_id) VALUES (?, ?)", newTaskName, todoListId)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}

	return nil
}
