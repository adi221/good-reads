package user

import (
	"github.com/adi221/good-reads/pkg/model"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/graphql-go/graphql"
)

var loginUserMutationField = &graphql.Field{
	Type:        userType,
	Description: "Log in a user",
	Args: graphql.FieldConfigArgument{
		"usernameOrEmail": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: loginResolver,
}

func loginResolver(p graphql.ResolveParams) (interface{}, error) {
	usernameOrEmail := p.Args["usernameOrEmail"].(string)
	password := p.Args["password"].(string)
	return service.Lookup().LoginUser(p.Context, usernameOrEmail, password)
}

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

	return service.Lookup().SignUpUser(p.Context, user)
}

func init() {
	schema.AddMutationField("signUpUser", signUpUserMutationField)
	schema.AddMutationField("loginUser", loginUserMutationField)
}
