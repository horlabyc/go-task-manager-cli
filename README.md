
# Go Task Manager

A command-line task management application written in Go. This application allows you to create, read, update, and delete tasks with support for filtering, sorting, and persistence.

## Features

* Task creation with title, description, and tags
* Task updating (title, description, completion status, tags)
* Task deletion
* Filtering tasks by completion status and tags
* Sorting tasks by creation date, update date, or title
* JSON file persistence
* Full command-line interface

## Installation

### Prerequisites

* Go 1.16 or higher

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/task-manager.git
cd task-manager

# Build the application
go build -o task-manager cmd/main.go
```

### Usage

```bash
# Add a new task
./task-manager add -title "Learn Go" -desc "Study Go programming" -tags "learning,programming"

# List all tasks
./task-manager list

# List only completed tasks
./task-manager list -completed

# List tasks filtered by tag
./task-manager list -tag "learning"

# List tasks sorted by title
./task-manager list -sort "title"

# Update a task (mark as completed)
./task-manager update -id 1 -completed

# Update a task's title and description
./task-manager update -id 1 -title "Master Go" -desc "Become proficient in Go programming"

# Delete a task
./task-manager delete -id 1
```

## Command Reference

### Add Task

```bash
./task-manager add -title "Task title" [-desc "Description"] [-tags "tag1,tag2"]
```

#### Options:

* `-title`: Task title (required)
* `-desc`: Task description (optional)
* `-tags`: Comma-separated list of tags (optional)

### Update Task

```bash
./task-manager update -id <task_id> [-title "New title"] [-desc "New description"] [-completed] [-tags "tag1,tag2"]
```

#### Options:

* `-id`: Task ID to update (required)
* `-title`: New task title (optional)
* `-desc`: New task description (optional)
* `-completed`: Mark task as completed (optional)
* `-tags`: New comma-separated list of tags (optional)

### Delete Task

```bash
./task-manager delete -id <task_id>
```

#### Options:

* `-id`: Task ID to delete (required)

### List Tasks

```bash
./task-manager list [-completed] [-tag "tagname"] [-sort "created|updated|title"]
```

#### Options:

* `-completed`: Filter to show only completed tasks (optional)
* `-tag`: Filter tasks by tag (optional)
* `-sort`: Sort tasks by "created", "updated", or "title" (optional)

## Project Structure

```bash
task-manager/
├── cmd/
│   └── main.go                # CLI entry point
├── internal/
│   └── task/
│       └── task.go            # Task manager implementation
├── go.mod                     # Module dependencies
└── README.md                  # This file
```
