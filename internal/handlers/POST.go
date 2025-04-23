package handlers

import (
	"FastAPI/internal/structures"
	"FastAPI/internal/variables"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func NewTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := uuid.New().String()
	newTask := &structures.Task{
		ID:     ID,
		Status: variables.StatusCreated,
	}

	variables.StorageMutex.Lock()
	variables.Storage[ID] = newTask
	variables.StorageMutex.Unlock()

	go func(task *structures.Task) {
		task.Status = variables.StatusRunning
		time.Sleep(time.Second * 5)
		task.Status = variables.StatusDone
	}(newTask)

	err := json.NewEncoder(w).Encode(map[string]string{
		"task_id": ID,
		"status":  "Created",
	})
	if err != nil {
		newTask.Error = "JSON encoding error"
		log.Printf("Error when encoding to JSON Task with task_id : %s", newTask.ID)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
}
