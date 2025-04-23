package executors

import (
	"FastAPI/internal/structures"
	"FastAPI/internal/variables"
	"math/rand"
	"time"
)

type SleepExecutor struct{}

func (s SleepExecutor) Execute(task *structures.Task) {
	task.Status = variables.StatusRunning
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := time.Duration(rnd.Intn(2)+3) * time.Minute
	time.Sleep(delay)
	task.Status = variables.StatusDone
}
