package executors

import "FastAPI/internal/structures"

type Executor interface {
	Execute(task *structures.Task)
}

var registry = make(map[string]Executor)

func Register(taskType string, executor Executor) {
	registry[taskType] = executor
}

func Get(taskType string) Executor {
	return registry[taskType]
}
