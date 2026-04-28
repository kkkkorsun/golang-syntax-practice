package main

import (
	"context"
	"fmt"

	"main.go/simple_connection"
	"main.go/simple_sql"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.Connect(ctx)
	if err != nil {
		panic(err)
	}

	createTableErr := simple_sql.CreateTable(ctx, conn)
	if createTableErr != nil {
		panic(createTableErr)
	}

	fmt.Println("Таблица успешно создана")

}
