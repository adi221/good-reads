package postgres

import (
	"database/sql"
	"strings"

	sq "github.com/Masterminds/squirrel"

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
	row := pg.db.QueryRow(query, args...)
	return mapRowToCategory(row)
}

func (pg *DB) GetCategoryByID(id uint) (*model.Category, error) {
	query, args, _ := pg.psql.Select(categoriesColumns...).From(
		categoriesTable,
	).Where(
		sq.Eq{"id": id},
	).ToSql()
	row := pg.db.QueryRow(query, args...)
	return mapRowToCategory(row)
}

func (pg *DB) GetCategoriesByUser(uid uint, filter model.FilterSchema) (*model.GetCategoriesResponse, error) {
	builder := pg.psql.Select(categoriesColumns...).From(
		categoriesTable,
	).Where(
		sq.Eq{"\"userId\"": uid},
	)

	// Maybe use count("*") because here we have to iterate each row again
	count, err := getCountFromQuery(pg, builder)
	if err != nil {
		return nil, err
	}

	offset := uint(0)
	if filter.Offset != nil {
		offset = *filter.Offset
	}

	limit := uint(20)
	if filter.Limit != nil {
		limit = *filter.Limit
	}

	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	// TODO: add sort by and sort order to query
	query, args, _ := builder.ToSql()
	rows, err := pg.db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category

	for rows.Next() {
		c := &model.Category{}
		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.Title,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &model.GetCategoriesResponse{
		Items: categories,
		Total: count,
	}, nil

}

func (pg *DB) CountCategoriesByUser(uid uint) (uint, error) {
	counter := pg.psql.Select("count(*)").From(
		categoriesTable,
	).Where(sq.Eq{"\"userId\"": uid})
	query, args, _ := counter.ToSql()

	var count uint
	if err := pg.db.QueryRow(query, args...).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
