package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"` // e.g., "Implement user authentication"
	Status      TaskStatus `json:"status"`      // e.g., "todo", "in-progress", "done", "archived"
	CreatedAt   string     `json:"created_at"`  // ISO 8601 format
	UpdatedAt   string     `json:"updated_at"`  // ISO 8601 format
}

func NewTask(id int, task string) *Task {
	return &Task{
		ID:          id,
		Description: task,
		Status:      StatusToDo,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("%+v\n", *t)
}
