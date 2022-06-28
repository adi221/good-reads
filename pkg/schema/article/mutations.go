package article

import (
	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/adi221/good-reads/pkg/util"
	"github.com/graphql-go/graphql"
)

var addArticleMutationField = &graphql.Field{
	Type:        articleType,
	Description: "Add a new article",
	Args: graphql.FieldConfigArgument{
		"url": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"categoryId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: addArticleResolver,
}

func addArticleResolver(p graphql.ResolveParams) (interface{}, error) {
	url, _ := p.Args["url"].(string)
	form := model.ArticleCreateForm{
		URL:        &url,
		CategoryID: util.GetGQLUintParameter(p.Args["categoryId"]),
	}

	return service.Lookup().CreateArticle(p.Context, form)
}

func init() {
	schema.AddMutationField("addArticle", addArticleMutationField)
}
