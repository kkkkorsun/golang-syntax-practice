package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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
	messages map[int64]Message
}

type Message struct {
	Title     string `json:"title"`
	Postcode  int64  `json:"postcode"`
	Body      string `json:"body"`
	IsExpress bool   `json:"is_express"`
}

func CreateEmptyMessage() *Message {
	return &Message{
		Title:     "",
		Postcode:  0,
		Body:      "",
		IsExpress: false,
	}
}

func NewMessageStorage() *MessageStorage {
	newMap := make(map[int64]Message)
	return &MessageStorage{sync.Mutex{}, newMap}
}

func (ms *MessageStorage) AddMessage(message Message) int64 {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	messageId := rand.Int63()
	ms.messages[messageId] = message

	return messageId
}
func (ms *MessageStorage) GetMessages() map[int64]Message {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	var messagesToSend = make(map[int64]Message)

	for k, v := range ms.messages {
		messagesToSend[k] = v
	}
	return messagesToSend
}

func (ms *MessageStorage) GetMessageById(id int64) (Message, bool) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	message, ok := ms.messages[id]
	return message, ok
}
func (ms *MessageStorage) DeleteMessage(id int64) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	delete(ms.messages, id)
}

func secondTask() {
	storage := NewMessageStorage()

	http.HandleFunc("/add-message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Ошибка, вызван неправильный метод"))
			return
		}

		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		userMessage := CreateEmptyMessage()
		parseErr := json.Unmarshal(httpRequestBody, &userMessage)
		if err != nil {
			fmt.Println(parseErr)
			return
		}

		messageId := storage.AddMessage(*userMessage)
		fmt.Println("Успешно добавили сообщение: ", userMessage, "с айди: ", messageId)

		fmt.Println("Текущие сообщения: ", storage.GetMessages())
	})

	http.HandleFunc("/delete-message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Ошибка, вызван неправильный метод"))
			return
		}

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

		message, ok := storage.GetMessageById(messageIdToDelete)
		if ok {
			fmt.Println("Удаляем сообщение: ", message)
			storage.DeleteMessage(messageIdToDelete)

			fmt.Println("Успешно удалили сообщение с айди: ", messageIdToDelete)

			fmt.Println("Текущие сообщения: ", storage.GetMessages())
		} else {
			fmt.Println("Сообщения для удаления с текущим айди не найдено")
		}
	})

	http.HandleFunc("/get-all-messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Ошибка, вызван неправильный метод"))
			return
		}

		messagesToSend, _ := json.Marshal(storage.GetMessages())

		w.Write(messagesToSend)

		fmt.Println(storage.GetMessages())
	})

	http.HandleFunc("/get-message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Ошибка, вызван неправильный метод"))
			return
		}

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
		messageIdToGet := int64(id)

		message, ok := storage.GetMessageById(messageIdToGet)
		if ok {
			w.WriteHeader(http.StatusOK)
			bytesToSend, _ := json.Marshal(message)
			_, err = w.Write(bytesToSend)
			if err != nil {
				return
			}

			fmt.Println("Запрошенное сообщение: ", message)
		} else {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte("Сообщения с текущим айди не найдено"))
			if err != nil {
				return
			}
		}
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

func thirdTask() {
	http.HandleFunc("/return-codes", func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		body := string(httpRequestBody)
		id, _ := strconv.Atoi(body)
		code := int64(id)

		switch code {
		case 200:
			w.WriteHeader(http.StatusOK)
			fmt.Println("Вернули 200")
		case 400:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Вернули 400")
		case 500:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Вернули 500")
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Этот код не обрабатывается, возвращаем 400")
		}
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

func fourthTask() {
	http.HandleFunc("/get-headers", func(w http.ResponseWriter, r *http.Request) {
		senderName := r.Header.Get("Sender-Name")
		if len(senderName) == 0 {
			fmt.Println("Передали пустое имя")
		} else {
			fmt.Println("Привет ", senderName)
		}
	})

	http.ListenAndServe(":8090", nil)
}

func fifthTask() {
	http.HandleFunc("/allowed-method", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Not Allowed"))
		} else {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("Allowed"))
		}
	})

	http.ListenAndServe(":8090", nil)
}
