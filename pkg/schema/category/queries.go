package category

import (
	"errors"
	"reflect"

	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/schema/common"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/adi221/good-reads/pkg/util"
	"github.com/graphql-go/graphql"
)

var categoriesQueryField = &graphql.Field{
	Type: categoriesResponseType,
	Args: graphql.FieldConfigArgument{
		"filter": &graphql.ArgumentConfig{
			Type: common.FilterSchemaInput,
		},
	},
	Resolve: categoriesResolver,
}

func categoriesResolver(p graphql.ResolveParams) (interface{}, error) {
	filter := p.Args["filter"]
	v := reflect.ValueOf(filter)
	var filterSchema model.FilterSchema
	switch v.Kind() {
	case reflect.Map:
		filterSchema = model.NewFilterSchema(filter.(map[string]interface{}))
	}
	return service.Lookup().GetCategories(p.Context, filterSchema)
}

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
	id, ok := util.ConvGQLParamaterToUint(p.Args["id"])
	if !ok {
		return nil, errors.New("invalid category ID")
	}
	return service.Lookup().GetCategory(p.Context, id)
}

func init() {
	schema.AddQueryField("category", categoryQueryField)
	schema.AddQueryField("categories", categoriesQueryField)
}
