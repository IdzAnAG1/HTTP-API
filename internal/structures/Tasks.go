package structures

type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Error  string `json:"error,omitempty"`
}
