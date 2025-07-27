package cli

import (
	"fmt"
	"os"

	"text/tabwriter"

	"github.com/jalvere00/task-tracker/internal/task"
)

func ListTasksHandler(taskList *task.TaskList) error {
	args := os.Args[2:]
	if len(args) < 1 {
		return listAllTasks(taskList)
	}

	switch args[0] {
	case "done":
		return listStatusTasks(taskList, task.StatusDone)
	case "todo":
		return listStatusTasks(taskList, task.StatusToDo)
	case "in-progress":
		return listStatusTasks(taskList, task.StatusInProgress)
	case "archive":
		return listStatusTasks(taskList, task.StatusArchived)
	default:
		return fmt.Errorf("there is no such list command as %s", args[0])
	}
}

func listAllTasks(taskList *task.TaskList) error {
	task := taskList.GetAllTasks()
	return printTask(task)
}

func listStatusTasks(taskList *task.TaskList, status task.TaskStatus) error {
	task := taskList.GetTasksByStatus(status)
	return printTask(task)
}

func printTask(tasks []*task.Task) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(writer, "ID\tStatus\tDescription")
	fmt.Fprintln(writer, "--\t------\t-----------")

	for _, t := range tasks {
		fmt.Fprintf(writer, "%d\t%s\t%s\n", t.ID, t.Status.String(), t.Description)
	}

	return writer.Flush()
}
