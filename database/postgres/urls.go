// Package postgres implements a postgres database functionalities.
package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"urlshort/types"
)

// CreateShort returns a created short link and an error.
func (p *psql) CreateShort(url string, length int) (string, error) {
	short := shorter(length)

	insertQuery := "INSERT INTO urls(origin, short) VALUES($1, $2)"

	_, err := p.db.Exec(insertQuery, url, short)
	if err != nil {
		return "", err
	}

	defer p.db.Close()

	return short, nil
}

// GetShort returns the given URL's short link and an error.
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

// GetOrigin returns the given short link's original URL and an error.
func (p *psql) GetOrigin(short string) (string, error) {
	var res string

	originQuery := "SELECT origin FROM urls WHERE short=$1"
	row := p.db.QueryRow(originQuery, short)
	err := row.Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("no origin url for: %s", short)
		}

		return "", err
	}

	return res, nil
}

// GetAll returns all URLs and their short links.
func (p *psql) GetAll() ([]types.URL, error) {
	url := []types.URL{}

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

		u := types.URL{
			Origin: origin,
			Short:  short,
		}

		url = append(url, u)
	}

	return url, nil
}

// shorter returns a random string with the given length.
func shorter(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	sid := rand.New(rand.NewSource(time.Now().UnixNano()))

	link := make([]byte, length)

	for k := range link {
		link[k] = charset[sid.Intn(len(charset))]
	}

	return string(link)
}
