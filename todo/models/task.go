package models

import "fmt"

var DefaultTaskList *TaskManager

type Task struct {
	ID    int64
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("标题为空")
	}
	return &Task{0, title, false}, nil
}

type TaskManager struct {
	tasks  []*Task
	lastID int64
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}

	return fmt.Errorf("未知的任务")
}

func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}

func (m *TaskManager) Find(ID int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}

func init() {
	DefaultTaskList = NewTaskManager()
}
