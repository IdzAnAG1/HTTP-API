package handlers

import (
	"FastAPI/internal/executors"
	"FastAPI/internal/structures"
	"FastAPI/internal/variables"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type NewTaskRequest struct {
	Type string `json:"type"`
}

func NewTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req NewTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "InvalidRequest"})
		log.Printf("Invalid Request")
		return
	}
	executor := executors.Get(req.Type)
	if executor == nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "UnknownTaskType"})
		log.Printf("Unknown Task type: %s", req.Type)
		return
	}
	ID := uuid.New().String()
	task := &structures.Task{
		ID:     ID,
		Type:   req.Type,
		Status: variables.StatusCreated,
	}
	variables.StorageMutex.Lock()
	variables.Storage[ID] = task
	variables.StorageMutex.Unlock()
	go executor.Execute(task)
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"task_id": ID,
		"status":  task.Status,
	})
}
