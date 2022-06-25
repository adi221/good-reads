package postgres

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

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

func (pg *DB) validateUserNotExistByUsername(username string) error {
	curUserByUsername, err := pg.GetUserByIdentity(username)
	if err != nil {
		return helper.ServerError()
	}
	if curUserByUsername != nil {
		return helper.NewHttpError(
			model.UsernameAlreadyExists,
			http.StatusConflict,
			"Username already exists",
		)
	}
	return nil
}

func (pg *DB) validateUserNotExistByEmail(email string) error {
	curUserByEmail, err := pg.GetUserByIdentity(email)
	if err != nil {
		return helper.ServerError()
	}
	if curUserByEmail != nil {
		return helper.NewHttpError(
			model.EmailAlreadyExists,
			http.StatusConflict,
			"Email already exists",
		)
	}
	return nil
}

func (pg *DB) CreateUser(user model.User) (*model.User, error) {
	if err := pg.validateUserNotExistByUsername(user.Username); err != nil {
		return nil, err
	}
	if err := pg.validateUserNotExistByEmail(user.Email); err != nil {
		return nil, err
	}

	hashedPassword, err := helper.ValidateAndEncryptPassword(user.Password)
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

	if verified := helper.VerifyPassword(result.Password, password); !verified {
		return nil, helper.NewHttpError(
			model.IncorrectCreds,
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
