package cli

import (
	"fmt"
	"os"

	"github.com/jalvere00/task-tracker/internal/task"
)

func AddTask(taskList *task.TaskList) error {
	args := os.Args[2:]
	if len(args) < 1 {
		return fmt.Errorf("description is required to add a task")
	}
	description := args[0]
	t := taskList.AddTask(description)
	fmt.Printf("Task added successfully (ID: %d)\n", t.ID)
	return nil
}
