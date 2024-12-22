package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var DBConnection *sql.DB

type HeroRequest struct {
	Powers []string `json:"powers"`
}

func SetUpRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/{todoListName}", CreateNewListHandler(DBConnection))
	r.Post("/add/{todoListName}", AddToListHandler())
	r.Get("/{todoListName}", GetListHandler())
	r.Delete("/{todoListName}", RemoveFromListHandler())
}

func setUpRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	SetUpRoutes(r)

	return r
}

func startServer(r *chi.Mux) {
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Error Starting Server")
	}
}

func Connect() *sql.DB {
	dsn := "user:password@tcp(localhost:3306)/db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}

func main() {
	DBConnection = Connect()
	defer DBConnection.Close()
	r := setUpRouter()
	startServer(r)
}
