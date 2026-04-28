package main

import (
	"restapi/http"
	"restapi/todo"
)

func main() {
	store := todo.CreateStore()
	handlers := http.NewHTTPHandlers(&store)
	server := http.NewHTTPServer(handlers)

	err := server.StartServer()
	if err != nil {
		return
	}
}
