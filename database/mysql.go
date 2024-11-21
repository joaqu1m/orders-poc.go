package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLCursor() (*sql.DB, error) {

	dsn := os.Getenv("MYSQL_DSN")

	return sql.Open("mysql", dsn)
}

func MySQLQuery(query string, args ...any) (*sql.Rows, error) {

	cursor, err := MySQLCursor()
	if err != nil {
		return nil, err
	}

	defer cursor.Close()
	return cursor.Query(query, args...)
}

func MySQLQueryRow(query string, args ...any) (*sql.Row, error) {

	cursor, err := MySQLCursor()
	if err != nil {
		return nil, err
	}

	defer cursor.Close()
	return cursor.QueryRow(query, args...), nil
}
