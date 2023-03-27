package database

import (
	"urlshort/types"
)

type DB interface {
	CreateShort(url string, length int) (string, error)
	GetShort(url string) (string, error)
	GetOrigin(short string) (string, error)
	GetAll() ([]types.URL, error)
}
