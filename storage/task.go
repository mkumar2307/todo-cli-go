package storage

// Task represents a single to-do item
type Task struct {
	Description string `json:"description"` // Task text
	Done        bool   `json:"done"`        // Completion status
}
