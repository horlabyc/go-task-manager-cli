package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/horlabyc/task-manager/internal/task"
)

func main() {
	manager, err := task.NewManager("tasks.json")
	if err != nil {
		log.Fatalf("Failed to create manager: %v", err)
	}

	// Define command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Task title")
	addDesc := addCmd.String("desc", "", "Task description")
	addTags := addCmd.String("tags", "", "Comma-separated tags")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "Task ID")
	updateTitle := updateCmd.String("title", "", "New task title")
	updateDesc := updateCmd.String("desc", "", "New task description")
	updateCompleted := updateCmd.Bool("completed", false, "Mark as completed")
	updateTags := updateCmd.String("tags", "", "New comma-separated tags")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "Task ID to delete")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCompleted := listCmd.Bool("completed", false, "Filter completed tasks")
	listTag := listCmd.String("tag", "", "Filter by tag")
	listSort := listCmd.String("sort", "", "Sort by: created, updated, or title")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" {
			log.Fatal("Title is required")
		}
		var tags []string
		if *addTags != "" {
			tags = strings.Split(*addTags, ",")
		}
		task, err := manager.CreateTask(*addTitle, *addDesc, tags)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created task: %+v\n", task)
	case "update":
		updateCmd.Parse(os.Args[2:])
		if *updateID == 0 {
			log.Fatal("Task ID is required")
		}

		var tags []string
		if *updateTags != "" {
			tags = strings.Split(*updateTags, ",")
		}

		task, err := manager.UpdateTask(*updateID, *updateTitle, *updateDesc, *updateCompleted, tags)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated task: %+v\n", task)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteID == 0 {
			log.Fatal("Task ID is required")
		}

		if err := manager.DeleteTask(*deleteID); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted task %d\n", *deleteID)
	case "list":
		listCmd.Parse(os.Args[2:])
		var completed *bool
		if listCmd.Lookup("completed").Value.String() != "false" {
			completed = listCompleted
		}

		tasks := manager.ListTasks(completed, *listTag, *listSort)
		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
			if t.Description != "" {
				fmt.Printf("   Description: %s\n", t.Description)
			}
			if len(t.Tags) > 0 {
				fmt.Printf("   Tags: %s\n", strings.Join(t.Tags, ", "))
			}
			fmt.Println()
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  task-manager add -title \"Task title\" [-desc \"Description\"] [-tags \"tag1,tag2\"]")
	fmt.Println("  task-manager update -id 1 [-title \"New title\"] [-desc \"New description\"] [-completed] [-tags \"tag1,tag2\"]")
	fmt.Println("  task-manager delete -id 1")
	fmt.Println("  task-manager list [-completed] [-tag \"tagname\"] [-sort \"created|updated|title\"]")
}
