package main

import (
	"log"

	"github.com/jalvere00/task-tracker/internal/app"
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
	a := app.NewApp()
	if err := a.Run(); err != nil {
		log.Fatalf("Error running the application: %v", err)
	}
}
