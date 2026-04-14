package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	secondTask()
}

func firstTask() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

type MessageStorage struct {
	mtx      sync.Mutex
	messages map[int64]string
}

func NewMessageStorage() *MessageStorage {
	newMap := make(map[int64]string)
	return &MessageStorage{sync.Mutex{}, newMap}
}

func (m *MessageStorage) AddMessage(id int64, message string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.messages[id] = message
}
func (m *MessageStorage) GetMessages() map[int64]string {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	var messagesToSend = make(map[int64]string)

	for k, v := range m.messages {
		messagesToSend[k] = v
	}
	return messagesToSend
}

func (m *MessageStorage) GetMessageById(id int64) string {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.messages[id]
}
func (m *MessageStorage) DeleteMessage(id int64) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	delete(m.messages, id)
}

func secondTask() {
	storage := NewMessageStorage()
	var messageId int64 = 0

	http.HandleFunc("/add-message", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		userMessage := string(httpRequestBody)

		storage.AddMessage(messageId, userMessage)
		fmt.Println("Успешно добавили сообщение: ", userMessage, "с айди: ", messageId)
		messageId++

		fmt.Println("Текущие сообщения: ", storage.GetMessages())
	})

	http.HandleFunc("/delete-message", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		body := string(httpRequestBody)
		id, err := strconv.Atoi(body)
		if err != nil {
			fmt.Println(err)
			return
		}

		messageIdToDelete := int64(id)
		fmt.Println("Удаляем сообщение: ", storage.GetMessageById(messageIdToDelete))
		storage.DeleteMessage(messageIdToDelete)

		fmt.Println("Успешно удалили сообщение с айди: ", messageIdToDelete)

		fmt.Println("Текущие сообщения: ", storage.GetMessages())
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
