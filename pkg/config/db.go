package config

import (
	"fmt"

	"github.com/adi221/good-reads/pkg/helper"
)

func GetDatabaseConnectionUri() string {
	user := helper.GetEnv("POSTGRES_CONN_USER", "postgres")
	password := helper.GetEnv("POSTGRES_CONN_PASSWORD", "postgres")
	host := helper.GetEnv("POSTGRESS_CONN_HOST", "localhost")
	port := helper.GetEnv("POSTGRESS_CONN_PORT", "5432")
	dbName := helper.GetEnv("POSTGRESS_CONN_NAME", "db")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
}
