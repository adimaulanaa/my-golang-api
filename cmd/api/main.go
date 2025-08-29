package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-go-api/internal/handlers"
	"my-go-api/internal/repository"
	"my-go-api/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Inisialisasi database
	connStr := "postgres://postgres:mypassword@localhost/my_go_api?sslmode=disable" // Ganti mypassword
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}
	defer db.Close()

	// Test koneksi
	if err := db.Ping(); err != nil {
		log.Fatal("Ping database gagal:", err)
	}
	log.Println("Berhasil konek ke database")

	// Inisialisasi repository dan service
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	handlers.SetTaskService(taskService)

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo, API Go pertamaku!")
	}).Methods("GET")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Start server
	fmt.Println("Server berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
