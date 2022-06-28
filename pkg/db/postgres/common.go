package postgres

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/adi221/good-reads/pkg/model"
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

func addFiltersToQuery(builder sq.SelectBuilder, filter model.FilterSchema) sq.SelectBuilder {
	offset := uint(0)
	if filter.Offset != nil {
		offset = *filter.Offset
	}

	limit := uint(20)
	if filter.Limit != nil {
		limit = *filter.Limit
	}

	sortBy := "\"createdAt\""
	if filter.SortBy != nil {
		sortBy = *filter.SortBy
	}

	sortOrder := "DESC"
	if filter.SortOrder != nil {
		sortOrder = *filter.SortOrder
	}

	return builder.Offset(uint64(offset)).Limit(uint64(limit)).OrderBy(sortBy + " " + strings.ToUpper(sortOrder))
}
