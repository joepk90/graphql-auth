package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type bodyRequest struct {
	Query     string
	Variables map[string]interface{}
}

// GQLHTTPMiddleware provides GraphQL as middleware
func GQLHTTPMiddleware(schema graphql.Schema) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body bodyRequest

		switch r.Method {
		case http.MethodGet:
			body = bodyRequest{
				Query: r.URL.Query().Get("query"),
			}
		default:
			defer r.Body.Close()
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
		}

		res := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  body.Query,
			RootObject:     make(map[string]interface{}),
			VariableValues: body.Variables,
			Context:        r.Context(),
		})

		json.NewEncoder(w).Encode(res)
	})
}
