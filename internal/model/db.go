package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres driver
)

type DB struct {
	db *sql.DB
}

func NewDB(username, password, host, dbname string) (*DB, error) {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		username, password, host, dbname,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, err
}
