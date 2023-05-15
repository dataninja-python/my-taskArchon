package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const tasksFile = "tasks.txt"

func loadTasks() []*Task {
	file, err := os.Open(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []*Task{}
		}
		fmt.Println("Error opening tasks file:", err)
		return nil
	}
	defer file.Close()

	var tasks []*Task
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		// Parse the task fields
		id, _ := strconv.Atoi(fields[0])
		desc := fields[1]
		createdAt, _ := time.Parse(time.RFC3339, fields[2])
		isCompleted, _ := strconv.ParseBool(fields[3])
		completedAt, _ := time.Parse(time.RFC3339, fields[4])
		pomodoroDone, _ := strconv.Atoi(fields[5])

		// Create a new task and add it to the slice
		task := &Task{
			ID:           id,
			Description:  desc,
			CreatedAt:    createdAt,
			CompletedAt:  completedAt,
			IsCompleted:  isCompleted,
			PomodoroDone: pomodoroDone,
		}
		tasks = append(tasks, task)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading tasks file:", err)
		return nil
	}

	return tasks
}

func saveTasks(tasks []*Task) error {
	file, err := os.Create(tasksFile)
	if err != nil {
		fmt.Println("Error creating tasks file:", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		line := fmt.Sprintf("%d|%s|%s|%t|%s|%d\n",
			task.ID, task.Description, task.CreatedAt.Format(time.RFC3339),
			task.IsCompleted, task.CompletedAt.Format(time.RFC3339), task.PomodoroDone)
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to tasks file:", err)
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing tasks file writer:", err)
		return err
	}

	return nil
}
