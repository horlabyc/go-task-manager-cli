package task

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Manager struct {
	tasks    []Task
	lastID   int
	filename string
}

// NewManager creates a new task manager with persistence
func NewManager(filename string) (*Manager, error) {
	m := &Manager{
		tasks:    make([]Task, 0),
		filename: filename,
	}

	// Load existing tasks if file exists
	if err := m.loadTasks(); err != nil {
		return nil, err
	}

	// Set last task ID to the highest ID in the list
	if len(m.tasks) > 0 {
		for _, task := range m.tasks {
			if task.ID > m.lastID {
				m.lastID = task.ID
			}
		}
	}
	return m, nil
}

func (m *Manager) loadTasks() error {
	data, err := os.ReadFile(m.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &m.tasks)
}

func (m *Manager) saveTasks() error {
	data, err := json.MarshalIndent(m.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.filename, data, 0644)
}

func (m *Manager) CreateTask(title, description string, tags []string) (*Task, error) {
	if title == "" {
		return nil, errors.New("task titlecannot be empty")
	}

	m.lastID++
	now := time.Now()

	task := Task{
		ID:          m.lastID,
		Title:       title,
		Description: description,
		Tags:        tags,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	m.tasks = append(m.tasks, task)
	if err := m.saveTasks(); err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *Manager) GetTaskByID(id int) (*Task, error) {
	for i, task := range m.tasks {
		if task.ID == id {
			return &m.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *Manager) UpdateTask(id int, title, description string, completed bool, tags []string) (*Task, error) {
	for i, task := range m.tasks {
		if task.ID == id {
			if title != "" {
				m.tasks[i].Title = title
			}
			if description != "" {
				m.tasks[i].Description = description
			}
			if tags != nil {
				m.tasks[i].Tags = tags
			}
			m.tasks[i].Completed = completed
			m.tasks[i].UpdatedAt = time.Now()
			if err := m.saveTasks(); err != nil {
				return nil, err
			}
			return &m.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *Manager) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return m.saveTasks()
		}
	}
	return errors.New("task not found")
}

func (m *Manager) ListTasks(filterCompleted *bool, filterTag string, sortBy string) []Task {
	tasks := make([]Task, len(m.tasks))
	copy(tasks, m.tasks)

	if filterCompleted != nil {
		filtered := make([]Task, 0)
		for _, task := range tasks {
			if task.Completed == *filterCompleted {
				filtered = append(filtered, task)
			}
		}
		tasks = filtered
	}

	if filterTag != "" {
		filtered := make([]Task, 0)
		for _, task := range tasks {
			for _, tag := range task.Tags {
				if strings.EqualFold(tag, filterTag) {
					filtered = append(filtered, task)
					break
				}
			}
		}
		tasks = filtered
	}

	switch strings.ToLower(sortBy) {
	case "created":
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
		})
	case "updated":
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].UpdatedAt.Before(tasks[j].UpdatedAt)
		})
	case "title":
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Title > tasks[j].Title
		})
	}
	return tasks
}
