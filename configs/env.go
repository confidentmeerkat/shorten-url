// Package configs store the program's configurations.
package configs

import "os"

var (
	Host     string = os.Getenv("POSTGRES_HOST")
	Port     string = os.Getenv("POSTGRES_PORT")
	User     string = os.Getenv("POSTGRES_USER")
	Password string = os.Getenv("POSTGRES_PASSWORD")
	DbName   string = os.Getenv("POSTGRES_DB")
	SSLMode  string = os.Getenv("POSTGRES_SSL_MODE")
	Domain   string = os.Getenv("SHORTENER_DOMAIN")
)
