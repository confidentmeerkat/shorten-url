package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"urlshort/database"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "postgres"
	password string = "toor"
	dbName   string = "shortener"
	sslMode  string = "disable"
)

type psql struct {
	db *sql.DB
}

func New() (database.DB, error) {
	db, err := sql.Open("postgres", psqlInfo())
	if err != nil {
		return nil, errors.New("failed to open the database")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("can not ping the database")
	}

	return &psql{db: db}, nil
}

func psqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)
}
