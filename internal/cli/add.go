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
	taskList.AddTask(description)
	return nil
}
