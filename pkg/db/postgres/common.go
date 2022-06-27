package postgres

import (
	sq "github.com/Masterminds/squirrel"
)

func getCountFromQuery(pg *DB, builder sq.SelectBuilder) (uint, error) {
	query, args, _ := builder.ToSql()
	rows, err := pg.db.Query(query, args...)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var counter uint
	for rows.Next() {
		counter++
	}
	err = rows.Err()
	if err != nil {
		return 0, err
	}

	return counter, nil
}
