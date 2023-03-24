package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "postgres"
	password string = ""
	dbName   string = "shortener"
	sslMode  string = "disable"
)

func psqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)
}

func Start() error {
	db, err := sql.Open("postgres", psqlInfo())
	if err != nil {
		return errors.New("failed to open the database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return errors.New("can not ping the database")
	}

	return nil
}
