package api

import (
	"encoding/json"
	"net/http"

	"context"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/schema"
	"github.com/adi221/good-reads/pkg/util"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/zerolog/log"

	// import all GraphQl schema
	_ "github.com/adi221/good-reads/pkg/schema/all"
)

var authExceptionOperations = map[string]interface{}{
	"loginUser":  true,
	"signUpUser": true,
}

func isOperationExemptFromAuth(operationName string) bool {
	_, ok := authExceptionOperations[operationName]
	return ok
}

func isRouteProtected(ctx context.Context) bool {
	uid := ctx.Value(constant.ContextUserID)
	return uid != nil
}

func graphqlHandler() http.Handler {
	root, err := schema.BuildRootSchema()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		opts := handler.NewRequestOptions(r)
		if !isRouteProtected(ctx) && !isOperationExemptFromAuth(opts.OperationName) {
			util.JsonErrors(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		params := graphql.Params{
			Schema:         root,
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			Context:        ctx,
		}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(result)
	})
}
