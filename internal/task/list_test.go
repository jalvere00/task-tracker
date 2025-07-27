package task

import (
	"testing"
)

func TestTaskList_AddTask(t *testing.T) {
	type testCase struct {
		description string
		expectedId  int
	}

	t.Run("AddTask with valid description", func(t *testing.T) {
		tl := NewTaskList()
		nextId := tl.GetNextId()

		tests := []testCase{
			{description: "First task", expectedId: 1},
			{description: "Second task", expectedId: 2},
			{description: "Third task", expectedId: 3},
		}
		for _, tc := range tests {
			task := tl.AddTask(tc.description)

			if task == nil {
				t.Fatalf("Expected a task to be added for description '%s', but got nil", tc.description)
			}

			if task.ID != nextId {
				t.Errorf("Expected task ID to be %d for description '%s', got %d", nextId, tc.description, task.ID)
			}

			if task.Description != tc.description {
				t.Errorf("Expected task description to be '%s' for ID %d, got '%s'", tc.description, nextId, task.Description)
			}

			if tl.GetNextId() != nextId+1 {
				t.Errorf("Expected next ID to be %d after adding task with description '%s', got %d", nextId+1, tc.description, tl.GetNextId())
			}
			nextId++
		}
	})
}

func TestTaskList_GetTaskById(t *testing.T) {
	tl := NewTaskList()
	tl.AddTask("First task")
	tl.AddTask("Second task")
	tl.AddTask("Third task")

	tests := []struct {
		id       int
		expected string
	}{
		{id: 1, expected: "First task"},
		{id: 2, expected: "Second task"},
		{id: 3, expected: "Third task"},
	}

	for _, tc := range tests {
		task, err := tl.GetTaskById(tc.id)

		if err != nil {
			t.Errorf("Expected no error for ID %d, but got: %v", tc.id, err)
			continue
		}
		if task == nil {
			t.Errorf("Expected task for ID %d, but got nil", tc.id)
			continue
		}

		if task.Description != tc.expected {
			t.Errorf("Expected task description to be '%s' for ID %d, got '%s'", tc.expected, tc.id, task.Description)
		}
	}
}

func TestTaskList_GetTaskList(t *testing.T) {
	tl := NewTaskList()
	tl.AddTask("ToDo 1 task")
	tl.AddTask("ToDo 2 task")

	task := tl.AddTask("Progress 1 task")
	tl.MarkTaskAsInProgress(task.ID)
	task = tl.AddTask("Progress 2 task")
	tl.MarkTaskAsInProgress(task.ID)

	task = tl.AddTask("Done task")
	tl.MarkTaskAsDone(task.ID)
	task = tl.AddTask("Done task")
	tl.MarkTaskAsDone(task.ID)

	t.Run("Get ToDo Tasks", func(t *testing.T) {
		todoList := tl.GetToDoTasks()
		if len(todoList) != 2 {
			t.Errorf("Not enough ToDO tasks")
		}

		for _, task := range todoList {
			if task.Status != StatusToDo {
				t.Errorf("The status of Task: %s doesn't have status 'todo'", task.Description)
			}
		}
	})

	t.Run("Get Progessing Tasks", func(t *testing.T) {
		todoList := tl.GetInProgressTasks()
		if len(todoList) != 2 {
			t.Errorf("Not enough In-Progress tasks")
		}

		for _, task := range todoList {
			if task.Status != StatusInProgress {
				t.Errorf("The status of Task: %s doesn't have status 'in-progress'", task.Description)
			}
		}
	})

	t.Run("Get Done Tasks", func(t *testing.T) {
		todoList := tl.GetDoneTasks()
		if len(todoList) != 2 {
			t.Errorf("Not enough Done tasks")
		}

		for _, task := range todoList {
			if task.Status != StatusDone {
				t.Errorf("The status of Task: %s doesn't have status 'done'", task.Description)
			}
		}
	})
}
