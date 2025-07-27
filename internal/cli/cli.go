package cli

import (
	"fmt"
	"os"

	"github.com/jalvere00/task-tracker/internal/json"
)

const outputFile = "output.json"

func HandleCommands() {
	taskList, err := json.ImportTaskListFromFile(outputFile)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	defer func() {
		err = json.ExportTaskListToFile(outputFile, taskList)
		if err != nil {
			panic(err) // Handle error appropriately in production code
		}
	}()

	args := os.Args

	switch args[1] {
	case "add":
		err = AddTask(taskList)
	case "update":
		err = UpdateTask(taskList)
	case "delete":
		err = DeleteTask(taskList)
	case "mark-in-progress":
		err = MarkTaskAsInProgress(taskList)
	case "mark-done":
		err = MarkTaskAsDone(taskList)
	case "list":
		err = ListTasksHandler(taskList)
	default:
		fmt.Println("Error")
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
