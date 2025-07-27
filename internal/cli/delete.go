package cli

import (
	"fmt"
	"os"

	"github.com/jalvere00/task-tracker/internal/task"
)

func DeleteTask(taskList *task.TaskList) error {
	args := os.Args[2:]
	if len(args) < 1 {
		return fmt.Errorf("delete command requires an taskID")
	}

	taskID, err := parseTaskID(args[0])
	if err != nil {
		return err
	}

	err = taskList.DeleteTask(taskID)
	if err != nil {
		return err
	}

	fmt.Printf("Task deleted task successfully (ID: %d)\n", taskID)
	return nil
}
