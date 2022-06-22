package article

import (
	"github.com/adi221/good-reads/pkg/helper"
	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
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
			Type: graphql.ID,
		},
	},
	Resolve: addArticleResolver,
}

func addArticleResolver(p graphql.ResolveParams) (interface{}, error) {
	var categoryId *uint
	if val, ok := helper.ConvGQLStringToUint(p.Args["categoryId"]); ok {
		categoryId = &val
	}
	url, _ := p.Args["url"].(string)
	form := model.ArticleCreateForm{
		URL:        &url,
		CategoryID: categoryId,
	}

	return service.Lookup().CreateArticle(p.Context, form)
}

func init() {
	schema.AddMutationField("addArticle", addArticleMutationField)
}
