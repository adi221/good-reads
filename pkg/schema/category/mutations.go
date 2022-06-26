package category

import (
	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/graphql-go/graphql"
)

var addCategoryMutationField = &graphql.Field{
	Type:        categoryType,
	Description: "Add a new category",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: addCategoryResolver,
}

func addCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	category := model.Category{
		Title: p.Args["title"].(string),
	}
	return service.Lookup().CreateCategory(p.Context, category)
}

func init() {
	schema.AddMutationField("addCategory", addCategoryMutationField)
}
