package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := loadTasks()

	for {
		fmt.Println("\n===== TaskWarrior Menu =====")
		fmt.Println("1. Add Task")
		fmt.Println("2. Complete Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Start Pomodoro")
		fmt.Println("5. Scrum Status")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			task := addTask(description)
			tasks = append(tasks, task)
			saveTasks(tasks)
			fmt.Println("Task added successfully.")
		case "2":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid task ID. Please try again.")
				continue
			}
			task := findTaskByID(tasks, taskID)
			if task == nil {
				fmt.Println("Task not found. Please try again.")
				continue
			}
			completeTask(task)
			saveTasks(tasks)
			fmt.Println("Task completed successfully.")
		case "3":
			listTasks(tasks)
		case "4":
			fmt.Print("Enter task ID to start Pomodoro: ")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid task ID. Please try again.")
				continue
			}
			task := findTaskByID(tasks, taskID)
			if task == nil {
				fmt.Println("Task not found. Please try again.")
				continue
			}
			startPomodoro(task)
			saveTasks(tasks)
		case "5":
			showScrumStatus()
		case "6":
			fmt.Println("Exiting TaskWarrior. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func listTasks(tasks []*Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func findTaskByID(tasks []*Task, id int) *Task {
	for _, task := range tasks {
		if task.ID == id {
			return task
		}
	}
	return nil
}
