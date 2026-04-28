package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Age       int64   `json:"age"`
	IsMarried bool    `json:"is_married"`
	Height    float64 `json:"height"`
}

func main() {
	http.HandleFunc("/send-user", func(writer http.ResponseWriter, request *http.Request) {
		userData := User{}

		httpRequestBody, err := io.ReadAll(request.Body)
		fmt.Println(string(httpRequestBody))
		if err != nil {
			fmt.Println(err)
			return
		}
		parseErr := json.Unmarshal(httpRequestBody, &userData)
		if parseErr != nil {
			fmt.Println("Ошибка при декодировании данных: ", err)
			return
		}

		fmt.Println(userData)
	})

	http.HandleFunc("/get-user", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			"Pavel",
			"Test",
			12,
			false,
			12.54}

		u, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		writer.Write(u)

	})

	http.ListenAndServe(":8080", nil)
}
