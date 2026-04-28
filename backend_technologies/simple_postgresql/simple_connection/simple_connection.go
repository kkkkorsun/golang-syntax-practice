package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	conn, connErr := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")
	if connErr != nil {
		return nil, connErr
	}

	fmt.Println("Подключение к базе прошло успешно")

	pingErr := conn.Ping(ctx)
	if pingErr != nil {
		return nil, pingErr
	}

	fmt.Println("Пинг в базу данныз отправлен успешно")

	return conn, nil
}
