package postgres

import (
	"database/sql"
	"fmt"
	"net/url"

	sq "github.com/Masterminds/squirrel"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// DB is a Database backed by PostgreSQL
type DB struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

// NewPostgreSQL creates a Database backed by PostgreSQL
func NewPostgreSQL(conn *url.URL) (*DB, error) {
	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		return nil, fmt.Errorf("Could not open PostgreSQL connection: %v", err)
	}

	// Test DB availablity
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Could not validate PostgreSQL connection: %v", err)
	}
	// Migrate DB if needed
	Migrate(db)

	return &DB{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil

}

// Close the DB.
func (pg *DB) Close() error {
	return pg.db.Close()
}
