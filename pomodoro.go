package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	pomodoroDuration   = 25 * time.Minute
	shortBreakDuration = 5 * time.Minute
	longBreakDuration  = 15 * time.Minute
)

func startPomodoro(task *Task) {
	fmt.Printf("Starting Pomodoro for Task #%d: %s\n", task.ID, task.Description)
	fmt.Println("Press 'Enter' to stop the Pomodoro.")

	timer := time.NewTimer(pomodoroDuration)

	go func() {
		<-timer.C
		fmt.Println("\nPomodoro completed!")
		task.PomodoroDone++

		if task.PomodoroDone%4 == 0 {
			startBreak(longBreakDuration)
		} else {
			startBreak(shortBreakDuration)
		}
	}()

	// Wait for user input to stop the Pomodoro
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	stopPomodoro(timer)
}

func startBreak(duration time.Duration) {
	fmt.Printf("\nStarting Break for %s\n", duration)
	timer := time.NewTimer(duration)
	<-timer.C
	fmt.Println("\nBreak ended!")
}

func stopPomodoro(timer *time.Timer) {
	if !timer.Stop() {
		<-timer.C
	}
	fmt.Println("Pomodoro stopped.")
}

func showScrumStatus() {
	tasks := loadTasks()

	totalTasks := len(tasks)
	completedTasks := 0
	pomodorosCompleted := 0

	for _, task := range tasks {
		if task.IsCompleted {
			completedTasks++
			pomodorosCompleted += task.PomodoroDone
		}
	}

	fmt.Println("Scrum Status:")
	fmt.Printf("Total Tasks: %d\n", totalTasks)
	fmt.Printf("Completed Tasks: %d\n", completedTasks)
	fmt.Printf("Pomodoros Completed: %d\n", pomodorosCompleted)
}
