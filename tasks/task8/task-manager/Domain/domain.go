package Domain

import "time"

// Task represents a task entity.
type Task struct {
    ID          string    `json:"id"` // Should be MongoDB ObjectID in hex format.
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"due_date"`
    Status      string    `json:"status"`
}

// User represents a user entity.
type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Password string `json:"password,omitempty"`
    Role     string `json:"role"` // "admin" or "user"
}