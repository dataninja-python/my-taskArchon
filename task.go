package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID           int
	Description  string
	CreatedAt    time.Time
	CompletedAt  time.Time
	IsCompleted  bool
	PomodoroDone int
}

func addTask(description string) *Task {
	task := &Task{
		ID:           generateTaskID(),
		Description:  description,
		CreatedAt:    time.Now(),
		IsCompleted:  false,
		PomodoroDone: 0,
	}

	return task
}

func completeTask(task *Task) {
	task.IsCompleted = true
	task.CompletedAt = time.Now()
}

func generateTaskID() int {
	tasks := loadTasks()

	if len(tasks) == 0 {
		return 1
	}

	lastTask := tasks[len(tasks)-1]
	return lastTask.ID + 1
}

func (task *Task) String() string {
	status := "Pending"
	if task.IsCompleted {
		status = "Completed"
	}

	return fmt.Sprintf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nCompleted At: %s\nPomodoros Completed: %d",
		task.ID, task.Description, status, task.CreatedAt.Format(time.RFC3339),
		task.CompletedAt.Format(time.RFC3339), task.PomodoroDone)
}
