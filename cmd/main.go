package main

import (
	"github.com/jalvere00/task-tracker/internal/cli"
)

// Task Tracker Application
// This application allows users to manage tasks with statuses like "todo", "in-progress", "done", and "archived".
// It provides functionalities to add, update, delete, and list tasks.

//  [Done] Create struct for Task
//  [Done] Create struct for TaskList
//  [ToDo] Create Test for TaskList
//  [ToDo] Implement CLI Command
//  [ToDo] Json File Handler

// cobra is not standard, use "flags" library instead
// us library "encoding/json"
func main() {
	cli.HandleCommands()
	// taskList := task.NewTaskList()
	// taskList.AddTask("Implement user authentication")
	// taskList.AddTask("Design database schema")
	// taskList.AddTask("Create REST API endpoints")
	// taskList.AddTask("Write unit tests for task management")

	// task1, err := taskList.GetTaskById(3)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Get Task by ID:", task1.String())

	// tl := task.NewTaskList()
	// tl.AddTask("ToDo 1 task")
	// tl.AddTask("ToDo 2 task")

	// task2 := tl.AddTask("Progress 1 task")
	// tl.MarkTaskAsInProgress(task2.ID)
	// task2 = tl.AddTask("Progress 2 task")
	// tl.MarkTaskAsInProgress(task2.ID)

	// task2 = tl.AddTask("Done task")
	// tl.MarkTaskAsDone(task2.ID)
	// task2 = tl.AddTask("Done task")
	// tl.MarkTaskAsDone(task2.ID)

	// json.ExportTaskListToFile("output.json", tl)
	// importTask, err := json.ImportTaskListFromFile("output.json")
	// if err != nil {
	// 	fmt.Println("Error importing task list:", err)
	// 	return
	// }
	// fmt.Printf("Imported Task List: %+v\n", importTask)
}
