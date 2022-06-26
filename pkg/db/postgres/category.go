package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/adi221/good-reads/pkg/model"
)

const categoriesTable = "categories"

var categoriesColumns = []string{
	"id",
	"\"userId\"",
	"title",
	"\"createdAt\"",
	"\"updatedAt\"",
}

func mapRowToCategory(row *sql.Row) (*model.Category, error) {
	c := &model.Category{}

	err := row.Scan(
		&c.ID,
		&c.UserID,
		&c.Title,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, mapError(err)
	}
	return c, nil
}

func (pg *DB) CreateCategoryForUser(uid uint, category model.Category) (*model.Category, error) {
	query, args, _ := pg.psql.Insert(
		categoriesTable,
	).Columns(
		"\"userId\"",
		"title",
	).Values(
		uid,
		category.Title,
	).Suffix(
		"RETURNING " + strings.Join(categoriesColumns, ","),
	).ToSql()
	fmt.Println(query)
	row := pg.db.QueryRow(query, args...)
	return mapRowToCategory(row)
}
