package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	taskToAdd := todo.CreateTask(taskDTO.Title, taskDTO.Description)
	err := h.todoList.AddTask(taskToAdd)
	if errors.Is(err, todo.ErrTaskAlreadyExists) {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusConflict)
		return
	} else if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(taskToAdd, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	t, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(t); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks := h.todoList.GetAllTasks()

	b, err := json.MarshalIndent(allTasks, "", "    ")
	if err != nil {
		panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		panic(err)
		return
	}
}

func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.GetUncompletedTasks()

	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		panic(err)
		return
	}
}
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	completeDTO := CompleteTaskDTO{}

	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	isComplete := completeDTO.Complete

	task, taskErr := h.todoList.SetTaskCompleted(title, isComplete)
	if taskErr != nil {
		errDTO := ErrorDTO{
			Message: taskErr.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	taskToSend, taskToSendErr := json.Marshal(task)
	if taskToSendErr != nil {
		panic(taskToSendErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(taskToSend)
	if err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}

}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	err := h.todoList.DeleteTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
