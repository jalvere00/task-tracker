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
		fmt.Println("Adding Command")
	case "update":
		fmt.Println("Update Command")
	case "delete":
		fmt.Println("Delete Command")
	case "mark-in-progress":
		fmt.Println("Marking Task as In-Progress")
	case "mark-done":
		fmt.Println("Marking Task as Done")
	case "list":
		fmt.Println("Listing Tasks")
	default:
		fmt.Println("Error")
	}

}
