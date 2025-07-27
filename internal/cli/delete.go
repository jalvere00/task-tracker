package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jalvere00/task-tracker/internal/task"
)

func DeleteTask(taskList *task.TaskList) error {
	args := os.Args[2:]
	if len(args) < 1 {
		return fmt.Errorf("delete command requires an taskID")
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = taskList.DeleteTask(taskID)
	if err != nil {
		return err
	}

	return nil
}
