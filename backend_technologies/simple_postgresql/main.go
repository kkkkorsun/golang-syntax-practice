package main

import (
	"context"

	"main.go/simple_connection"
)

func main() {
	ctx := context.Background()

	err := simple_connection.Connect(ctx)
	if err != nil {
		panic(err)
	}
}
