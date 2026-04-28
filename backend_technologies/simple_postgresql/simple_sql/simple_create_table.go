package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//id, name (not null), author(not null), review, year of release (not null), is fully read bool, time when book fully read

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := ` CREATE TABLE IF NOT EXISTS books(
     id SERIAL PRIMARY KEY,
     author VARCHAR(50) NOT NULL, 
     review VARCHAR(100), 
     year_of_release DATE NOT NULL, 
     is_read BOOLEAN NOT NULL, 
     time_when_book_was_read TIMESTAMP
 )`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
