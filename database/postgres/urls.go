package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"urlshort/types"
)

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
	var short string

	query := "SELECT short FROM urls WHERE origin=$1"
	row := p.db.QueryRow(query, url)
	err := row.Scan(&short)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("there are no short link for the given url")
		}

		return "", err
	}

	return short, nil
}

func (p *psql) GetOrigin(short string) (string, error) {
	return "", nil
}

func (p *psql) GetAll() ([]types.Url, error) {
	url := []types.Url{}

	allQuery := "SELECT origin, short FROM urls"

	rows, err := p.db.Query(allQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		origin, short := "", ""

		err = rows.Scan(&origin, &short)
		if err != nil {
			return nil, err
		}

		u := types.Url{
			Origin: origin,
			Short:  short,
		}

		url = append(url, u)
	}

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
