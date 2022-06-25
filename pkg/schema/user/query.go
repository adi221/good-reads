package user

import (
	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/graphql-go/graphql"
)

var meQueryField = &graphql.Field{
	Type:        userType,
	Description: "Get current user",
	Resolve:     meResolver,
}

func meResolver(p graphql.ResolveParams) (interface{}, error) {
	userFromContext := p.Context.Value(constant.ContextUser)
	return userFromContext, nil
}

func init() {
	schema.AddQueryField("me", meQueryField)
}
