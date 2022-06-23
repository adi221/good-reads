package postgres

import (
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/adi221/good-reads/pkg/model"
)

const usersTable = "users"

var usersColumns = []string{
	"id",
	"username",
	"email",
	"createdAt",
}

func mapRowToUser(row *sql.Row) (*model.User, error) {
	user := &model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validateAndEncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return password, err
	}
	return string(hashedPassword), nil
}

func (pg *DB) CreateUser(user model.User) (*model.User, error) {
	hashedPassword, err := validateAndEncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}

	query, args, _ := pg.psql.Insert(
		usersTable,
	).Columns(
		"username",
		"email",
		"password",
	).Values(
		user.Username,
		user.Email,
		hashedPassword,
	).Suffix(
		"RETURNING " + strings.Join(articleColumns, ","),
	).ToSql()
	row := pg.db.QueryRow(query, args...)
	return mapRowToUser(row)
}
