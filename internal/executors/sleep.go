package executors

import (
	"FastAPI/internal/structures"
	"FastAPI/internal/variables"
	"time"
)

type SleepExecutor struct{}

func (s SleepExecutor) Execute(task *structures.Task) {
	task.Status = variables.StatusRunning
	time.Sleep(5 * time.Minute)
	task.Status = variables.StatusDone
}
