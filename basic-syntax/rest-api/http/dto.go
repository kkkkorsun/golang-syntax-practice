package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type CompleteTaskDTO struct {
	Complete bool `json:"complete"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
