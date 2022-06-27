package common

import "github.com/graphql-go/graphql"

// FilterSchemaInput is a input to get-bulk requests with selection paramaters
var FilterSchemaInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "FilterSchemaInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"limit": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"sortBy": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"sortOrder": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
