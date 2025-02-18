package task

import (
	"errors"
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
	tasks  []Task
	lastID int
}

func NewManager() *Manager {
	return &Manager{
		tasks:  make([]Task, 0),
		lastID: 0,
	}
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
