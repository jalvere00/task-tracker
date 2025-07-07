package task

import (
	"fmt"
	"time"
)

type TaskList struct {
	Tasks  []*Task `json:"tasks"`
	NextId int     `json:"next_id"` // The next ID to be assigned to a new task
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  []*Task{},
		NextId: 1,
	}
}

func (tl *TaskList) GetNextId() int {
	return tl.NextId
}

func (tl *TaskList) AddTask(description string) *Task {
	task := NewTask(tl.NextId, description)
	tl.Tasks = append(tl.Tasks, task)
	tl.NextId++
	return task
}

func (tl *TaskList) GetAllTasks() []*Task {
	return tl.Tasks
}

func (tl *TaskList) GetTasksByStatus(status TaskStatus) []*Task {
	var filteredTasks []*Task
	for _, task := range tl.Tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func (tl *TaskList) GetToDoTasks() []*Task {
	return tl.GetTasksByStatus(StatusToDo)
}

func (tl *TaskList) GetInProgressTasks() []*Task {
	return tl.GetTasksByStatus(StatusInProgress)
}

func (tl *TaskList) GetDoneTasks() []*Task {
	return tl.GetTasksByStatus(StatusDone)
}

func (tl *TaskList) GetTaskById(id int) (*Task, error) {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) DeleteTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) UpdateTask(id int, description string) error {
	task, err := tl.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}
	task.Description = description
	task.UpdatedAt = time.Now().Format(time.RFC3339)
	return nil
}

func (tl *TaskList) updateTaskStatus(id int, status TaskStatus) error {
	task, err := tl.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("error updating task status: %w", err)
	}
	task.Status = status
	task.UpdatedAt = time.Now().Format(time.RFC3339)
	return nil
}

func (tl *TaskList) MarkTaskAsDone(id int) error {
	return tl.updateTaskStatus(id, StatusDone)
}

func (tl *TaskList) MarkTaskAsInProgress(id int) error {
	return tl.updateTaskStatus(id, StatusInProgress)
}

func (tl *TaskList) MarkTaskAsArchived(id int) error {
	return tl.updateTaskStatus(id, StatusArchived)
}
