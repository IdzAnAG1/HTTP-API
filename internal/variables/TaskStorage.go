package variables

import (
	"FastAPI/internal/structures"
	"sync"
)

var (
	Storage      = make(map[string]*structures.Task)
	StorageMutex = sync.Mutex{}
)
