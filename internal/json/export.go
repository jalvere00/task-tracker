package json

import (
	"encoding/json"
	"os"

	"github.com/jalvere00/task-tracker/internal/task"
)

// ExportTaskListToFile writes the TaskList to the specified file in JSON format.
func ExportTaskListToFile(filename string, taskList *task.TaskList) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(taskList)
}

// ImportTaskListFromFile reads a TaskList from the specified file in JSON format.
func ImportTaskListFromFile(filename string) (*task.TaskList, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return task.NewTaskList(), nil // Return a new TaskList if the file does not exist
		}
		return nil, err
	}
	defer file.Close()

	var taskList task.TaskList
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&taskList); err != nil {
		return nil, err
	}

	return &taskList, nil
}
