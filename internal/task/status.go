package task

import "fmt"

type TaskStatus int

const (
	StatusToDo TaskStatus = iota
	StatusInProgress
	StatusDone
	StatusArchived
)

func (s TaskStatus) String() string {
	switch s {
	case StatusToDo:
		return "todo"
	case StatusInProgress:
		return "in-progress"
	case StatusDone:
		return "done"
	case StatusArchived:
		return "archived"
	default:
		return "unknown"
	}
}

func ParseTaskStatus(status string) (TaskStatus, error) {
	switch status {
	case "todo":
		return StatusToDo, nil
	case "in-progress":
		return StatusInProgress, nil
	case "done":
		return StatusDone, nil
	case "archived":
		return StatusArchived, nil
	default:
		return -1, fmt.Errorf("invalid task status: %s", status)
	}
}
