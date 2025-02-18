package task

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
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

func (m *Manager) CreateTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("task titlecannot be empty")
	}

	m.lastID++
	now := time.Now()

	task := Task{
		ID:          m.lastID,
		Title:       title,
		Description: description,
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
