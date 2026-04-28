package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) error {
	conn, connErr := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")
	if connErr != nil {
		return connErr
	}

	fmt.Println("Подключение к базе прошло успешно")

	pingErr := conn.Ping(ctx)
	if pingErr != nil {
		return pingErr
	}

	fmt.Println("Пинг в базу данныз отправлен успешно")

	return nil
}
