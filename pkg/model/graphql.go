package model

import "github.com/adi221/good-reads/pkg/util"

type FilterSchema struct {
	Limit     *uint
	Offset    *uint
	SortOrder *string
	SortBy    *string
}

func NewFilterSchema(filter map[string]interface{}) FilterSchema {
	return FilterSchema{
		Limit:     util.GetGQLUintParameter(filter["limit"]),
		Offset:    util.GetGQLUintParameter(filter["offset"]),
		SortBy:    util.GetGQLStringParameter(filter["sortBy"]),
		SortOrder: util.GetGQLStringParameter(filter["sortOrder"]),
	}
}
