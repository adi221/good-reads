package article

import (
	"errors"
	"fmt"

	"github.com/adi221/good-reads/pkg/schema"
	"github.com/graphql-go/graphql"
)

var articleQueryField = &graphql.Field{
	Type: articleType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: articleResolver,
}

type TempArticle struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func articleResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println(p.Args)
	id, ok := p.Args["id"]
	if !ok {
		return nil, errors.New("ID is missing")
	}
	// here we should search for an article with id, but for now:
	a := &TempArticle{id.(string), "Good job!"}
	return a, nil
}

func init() {
	schema.AddQueryField("article", articleQueryField)
}
