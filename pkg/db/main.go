package db

import (
	"fmt"
	"net/url"

	"github.com/adi221/good-reads/pkg/db/postgres"
	"github.com/rs/zerolog/log"
)

// DB is the global database structure
type DB interface {
	Close() error
	ArticleRepository
	UserRespository
	CategoryRespository
}

// NewDB creates a new database provider regarding the datasource uri
func NewDB(conn string) (DB, error) {
	u, err := url.ParseRequestURI(conn)
	if err != nil {
		return nil, fmt.Errorf("Invalid conntection URL: %s", conn)
	}
	provider := u.Scheme
	var db DB

	switch provider {
	case "postgres":
		db, err = postgres.NewPostgreSQL(u)
		if err != nil {
			return nil, err
		}
		log.Info().Str("component", "database").Str("uri", u.Redacted()).Msg("using PostgreSQL database")
	}

	return db, nil
}
