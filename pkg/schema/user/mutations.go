package user

import (
	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/graphql-go/graphql"
)

var signUpUserMutationField = &graphql.Field{
	Type:        userType,
	Description: "Sign up a user",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: signUpResolver,
}

func signUpResolver(p graphql.ResolveParams) (interface{}, error) {
	user := model.User{
		Username: p.Args["username"].(string),
		Email:    p.Args["email"].(string),
		Password: p.Args["password"].(string),
	}

	return service.Lookup().SignUpUser(user)
}

func init() {
	schema.AddMutationField("signUpUser", signUpUserMutationField)
}
