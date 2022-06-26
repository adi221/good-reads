package category

import (
	"github.com/graphql-go/graphql"
)

var categoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"userId": &graphql.Field{
				Type: graphql.ID,
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
