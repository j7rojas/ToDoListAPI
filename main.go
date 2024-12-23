package main

import (
	"ToDoListAPI/config"
	"ToDoListAPI/routes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func startServer(r *chi.Mux) {
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Error Starting Server")
	}
}

func main() {
	DBConnection := config.Connect()
	defer DBConnection.Close()

	r := routes.SetUpRouter(DBConnection)
	startServer(r)
}
