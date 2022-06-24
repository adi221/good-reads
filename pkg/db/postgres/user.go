package postgres

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/adi221/good-reads/pkg/helper"
	"github.com/adi221/good-reads/pkg/model"
)

const usersTable = "users"

var usersColumns = []string{
	"id",
	"username",
	"email",
	"\"createdAt\"",
	"\"updatedAt\"",
}

var userColumnsWithEncryptedPassword = []string{
	"id",
	"username",
	"email",
	"password",
	"\"createdAt\"",
	"\"updatedAt\"",
}

func mapRowToUser(row *sql.Row) (*model.User, error) {
	user := &model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func mapRowToUserWithPassword(row *sql.Row) (*model.User, error) {
	user := &model.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
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

func verifyPassword(encryptedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		return false
	}
	return true
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
		"RETURNING " + strings.Join(usersColumns, ","),
	).ToSql()
	row := pg.db.QueryRow(query, args...)
	return mapRowToUser(row)
}

func (pg *DB) GetUserByIdentityAndVerify(identity string, password string) (*model.User, error) {
	row := pg.db.QueryRow(
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE username = $1 OR email = $1",
			strings.Join(userColumnsWithEncryptedPassword, ","),
			usersTable,
		),
		identity,
	)

	result, err := mapRowToUserWithPassword(row)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if verified := verifyPassword(result.Password, password); !verified {
		return nil, helper.NewHttpError(
			model.ErrIncorrectCreds,
			http.StatusUnauthorized,
			"Password is incorrect",
		)
	} else {
		result.Password = ""
	}

	return result, nil
}

func (pg *DB) GetUserByIdentity(identity string) (*model.User, error) {
	row := pg.db.QueryRow(
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE username = $1 OR email = $1",
			strings.Join(usersColumns, ","),
			usersTable,
		),
		identity,
	)

	result, err := mapRowToUser(row)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
