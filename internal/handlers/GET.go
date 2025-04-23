package handlers

import (
	"FastAPI/internal/variables"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ShowTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskID := mux.Vars(r)["id"]

	task, exist := variables.Storage[taskID]
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "TaskNotFound",
		})
		if err != nil {
			log.Printf("Task with task_id: %s Not Found", taskID)
		}
		return
	}

	response := map[string]interface{}{
		"task_id": task.ID,
		"status":  task.Status,
	}
	if task.Error != "" {
		response["error"] = task.Error
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error when encoding response: %v", err)
	}
}
func ShowAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []map[string]interface{}
	for _, task := range variables.Storage {
		taskResponse := map[string]interface{}{
			"task_id": task.ID,
			"status":  task.Status,
		}

		if task.Error != "" {
			taskResponse["error"] = task.Error
		}

		tasks = append(tasks, taskResponse)
	}
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "InternalServerError",
		})
	}
}
