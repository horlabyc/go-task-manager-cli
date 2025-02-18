package main

import (
	"fmt"
	"log"

	"github.com/horlabyc/task-manager/internal/task"
)

func main() {
	manager := task.NewManager()

	task1, err := manager.CreateTask("Buy groceries", "Buy groceries for the week")
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	task2, err := manager.CreateTask("Learn Go Basics", "Study fundamental Go concepts")
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	fmt.Printf("Task 1: %v\n", task1)
	fmt.Printf("Task 2: %v\n", task2)

	// Retrieve and display tasks
	task, err := manager.GetTaskByID(task1.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve task: %v", err)
	}

	fmt.Printf("Retrieved task: %v\n", task)
}
