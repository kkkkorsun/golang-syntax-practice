package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func CreateTask(taskTitle string, taskDescription string) Task {
	return Task{
		Title:       taskTitle,
		Description: taskDescription,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

func (t *Task) UpdateTaskStatus(isCompleted bool) {
	if isCompleted {
		t.Completed = true
		completedTime := time.Now()
		t.CompletedAt = &completedTime
	} else {
		t.Completed = false
		t.CompletedAt = nil
	}
}

func (t *Task) Complete() {
	t.Completed = true
	completedTime := time.Now()
	t.CompletedAt = &completedTime
}

func (t *Task) UnComplete() {
	t.Completed = false
	t.CompletedAt = nil
}
