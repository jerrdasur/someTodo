package todo

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type Task struct {
	ID    string
	Title string
}

type List struct {
	mx    sync.RWMutex
	todos []Task
}

func NewList() *List {
	return &List{}
}

func (l *List) List() []Task {
	l.mx.RLock()
	defer l.mx.RUnlock()

	return l.todos
}

func (l *List) Get(todoID string) *Task {
	l.mx.RLock()
	defer l.mx.RUnlock()

	for _, t := range l.todos {
		if t.ID == todoID {
			return &t
		}
	}

	return nil
}

func (l *List) Add(title string) Task {
	l.mx.Lock()
	l.mx.Unlock()

	t := Task{
		ID:    uuid.NewString(),
		Title: title,
	}
	l.todos = append(l.todos, t)

	return t
}

func (l *List) Remove(todoID string) error {
	l.mx.Lock()
	defer l.mx.Unlock()

	newList := make([]Task, 0, len(l.todos))

	var found bool
	for _, t := range l.todos {
		if t.ID == todoID {
			found = true
			continue
		}

		newList = append(newList, t)
	}

	if !found {
		return ErrTodoNotFound
	}

	l.todos = newList

	return nil
}

func (l *List) Update(t Task) error {
	l.mx.Lock()
	defer l.mx.Unlock()

	var found bool

	for i := range l.todos {
		if l.todos[i].ID == t.ID {
			found = true
			l.todos[i] = t
		}
	}

	if !found {
		return ErrTodoNotFound
	}

	return nil
}
