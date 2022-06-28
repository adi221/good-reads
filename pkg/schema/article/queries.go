package article

import (
	"errors"

	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/adi221/good-reads/pkg/util"
	"github.com/graphql-go/graphql"
)

var articleQueryField = &graphql.Field{
	Type: articleType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: articleResolver,
}

func articleResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := util.ConvGQLParamaterToUint(p.Args["id"])
	if !ok {
		return nil, errors.New("invalid article ID")
	}
	return service.Lookup().GetArticle(p.Context, id)
}

func init() {
	schema.AddQueryField("article", articleQueryField)
}
