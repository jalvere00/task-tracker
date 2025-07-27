package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jalvere00/task-tracker/internal/task"
)

func UpdateTask(taskList *task.TaskList) error {
	args := os.Args[2:]
	if len(args) < 2 {
		return fmt.Errorf("task ID and new description are required to update a task")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	newDescription := args[1]
	t, err := taskList.GetTaskById(taskID)
	if err != nil {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	t.Description = newDescription
	return nil
}
