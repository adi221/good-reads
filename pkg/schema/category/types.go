package category

import (
	"github.com/graphql-go/graphql"
)

var categoriesResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CategoriesResponse",
		Fields: graphql.Fields{
			"items": &graphql.Field{
				Type: graphql.NewList(categoryType),
			},
			"total": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var categoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
