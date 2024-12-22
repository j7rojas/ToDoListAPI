package main

import (
	"database/sql"
	"fmt"
)

type TodoList struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	ITEMS    []Item `json:"items"`
	COMPLETE bool   `json:"complete"`
}

type Item struct {
	NAME     string `json:"name"`
	COMPLETE bool   `json:"complete"`
}

func InsertTodo(db *sql.DB, name string, complete bool) error {
	_, err := db.Exec("INSERT INTO todos (`name`, `complete`) VALUES (?, ?)", name, complete)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	return nil
}
