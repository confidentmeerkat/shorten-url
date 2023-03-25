package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"urlshort/database"
	"urlshort/types"

	_ "github.com/lib/pq"
)

type psql struct {
	db *sql.DB
}

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "postgres"
	password string = "toor"
	dbName   string = "shortener"
	sslMode  string = "disable"
)

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

func (p *psql) CreateShort(url string, length int) (string, error) {
	short := newShort(length)

	insertQuery := "INSERT INTO urls(origin, short) VALUES($1, $2)"

	_, err := p.db.Exec(insertQuery, url, short)
	if err != nil {
		return "", err
	}

	defer p.db.Close()

	return short, nil
}

func (p *psql) GetShort(url string) (string, error) {
	return "", nil
}

func (p *psql) GetOrigin(short string) (string, error) {
	return "", nil
}

func (p *psql) GetAll() ([]types.Url, error) {
	url := []types.Url{}

	return url, nil
}

func psqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)
}

func newShort(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	sid := rand.New(rand.NewSource(time.Now().UnixNano()))

	link := make([]byte, length)

	for k := range link {
		link[k] = charset[sid.Intn(len(charset))]
	}

	return string(link)

}
