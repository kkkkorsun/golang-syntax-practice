package todo

import (
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func CreateStore() List {
	return List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[task.Title]
	if ok {
		return ErrTaskAlreadyExists
	}

	l.tasks[task.Title] = task
	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if ok {
		return task, nil
	}

	return Task{}, ErrTaskNotFound
}

func (l *List) GetAllTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tempMap := make(map[string]Task)

	for k, v := range l.tasks {
		tempMap[k] = v
	}

	return tempMap
}

func (l *List) GetUncompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tempMap := make(map[string]Task)

	for k, v := range l.tasks {
		if !v.Completed {
			tempMap[k] = v
		}
	}

	return tempMap
}

func (l *List) GetCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tempMap := make(map[string]Task)

	for k, v := range l.tasks {
		if v.Completed {
			tempMap[k] = v
		}
	}

	return tempMap
}

func (l *List) SetTaskCompleted(title string, isCompleted bool) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if ok {
		task.UpdateTaskStatus(isCompleted)
		l.tasks[title] = task
		return task, nil
	}

	return Task{}, ErrTaskNotFound
}

func (l *List) UnCompleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if ok {
		task.UnComplete()
		l.tasks[title] = task
		return nil
	}

	return ErrTaskNotFound
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if ok {
		delete(l.tasks, task.Title)
		return nil
	}

	return ErrTaskNotFound
}
