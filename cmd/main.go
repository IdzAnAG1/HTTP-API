package main

import (
	"FastAPI/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/task", handlers.NewTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", handlers.ShowAllTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", handlers.ShowTaskHandler).Methods("GET")

	port := ":8080"
	log.Printf("Server is running on %s", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
