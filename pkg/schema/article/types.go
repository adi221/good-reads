package article

import (
	"github.com/graphql-go/graphql"
)

var articleStatus = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "status",
		Description: "Article status",
		Values: graphql.EnumValueConfigMap{
			"read": &graphql.EnumValueConfig{
				Value:       "read",
				Description: "article is read",
			},
			"unread": &graphql.EnumValueConfig{
				Value:       "unread",
				Description: "article is not read yet",
			},
		},
	},
)

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"html": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: articleStatus,
			},
			// TODO: change to category resolver once we have getCategory handler
			"categoryId": &graphql.Field{
				Type: graphql.Int,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
