package config

import (
	"fmt"

	"github.com/adi221/good-reads/pkg/util"
)

func GetDatabaseConnectionUri() string {
	user := util.GetEnv("POSTGRES_CONN_USER", "postgres")
	password := util.GetEnv("POSTGRES_CONN_PASSWORD", "postgres")
	host := util.GetEnv("POSTGRESS_CONN_HOST", "localhost")
	port := util.GetEnv("POSTGRESS_CONN_PORT", "5432")
	dbName := util.GetEnv("POSTGRESS_CONN_NAME", "db")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
}
