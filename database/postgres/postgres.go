package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"urlshort/database"

	_ "github.com/lib/pq"
)

var (
	host     string = os.Getenv("POSTGRES_HOST")
	port     string = os.Getenv("POSTGRES_PORT")
	user     string = os.Getenv("POSTGRES_USER")
	password string = os.Getenv("POSTGRES_PASSWORD")
	dbName   string = os.Getenv("POSTGRES_DB_NAME")
	sslMode  string = os.Getenv("POSTGRES_SSL_MODE")
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
