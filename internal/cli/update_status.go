package cli

import (
	"fmt"
	"os"

	"github.com/jalvere00/task-tracker/internal/task"
)

type TaskStatus func(int) error

func MarkTaskAsDone(taskList *task.TaskList) error {
	return markTaskStatus(taskList.MarkTaskAsDone)
}

func MarkTaskAsInProgress(taskList *task.TaskList) error {
	return markTaskStatus(taskList.MarkTaskAsInProgress)
}

func MarkTaskAsArchived(taskList *task.TaskList) error {
	return markTaskStatus(taskList.MarkTaskAsArchived)
}

func markTaskStatus(operator TaskStatus) error {
	args := os.Args[2:]
	if len(args) < 1 {
		return fmt.Errorf("delete command requires an taskID")
	}

	taskId, err := parseTaskID(args[0])
	if err != nil {
		return err
	}
	return operator(taskId)
}
