package category

import (
	"errors"

	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/adi221/good-reads/pkg/util"
	"github.com/graphql-go/graphql"
)

var categoryQueryField = &graphql.Field{
	Type: categoryType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: categoryResolver,
}

func categoryResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := util.ConvGQLStringToUint(p.Args["id"])
	if !ok {
		return nil, errors.New("invalid category ID")
	}
	return service.Lookup().GetCategory(p.Context, id)
}

func init() {
	schema.AddQueryField("category", categoryQueryField)
}
