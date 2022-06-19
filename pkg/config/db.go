package config

import (
	"fmt"
	"os"
)

func GetDatabaseConnectionUri() string {
	// TODO: prefer to use the whole connection string as a env var, but the special characters in the string should be escaped so skip it for now.
	password := os.Getenv("POSTGRES_CONN_PASSWORD")
	if password == "" {
		panic("'POSTGRES_CONN_PASSWORD' environment variable must be set")
	}
	return fmt.Sprintf("postgres://postgres:%s@localhost:5433/good-reads?sslmode=disable", password)
}
