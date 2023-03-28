// Package postgres implements a postgres database functionalities.
package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"urlshort/configs"
	"urlshort/database"

	_ "github.com/lib/pq"
)

type psql struct {
	db *sql.DB
}

// psqlInfo returns configuration for connecting to a postgres database.
func psqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Host, configs.Port, configs.User, configs.Password, configs.DbName, configs.SSLMode,
	)
}

// New creates a new connection to postgres database.
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
