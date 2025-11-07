package policy

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/joepk90/graphql-auth/internal/auth"
	"github.com/joepk90/graphql-auth/internal/stats"
)

//go:generate go run go.uber.org/mock/mockgen -package=mocks -destination=../mocks/metrics.go github.com/joepk90/graphql-auth/internal/stats Metrics
//go:generate go run go.uber.org/mock/mockgen -package=mocks -destination=../mocks/authorizer_mock.go github.com/joepk90/graphql-auth/internal/auth AuthorizerInterface

const (
	queryType         = "query"
	mutationQueryType = "mutation"
)

// Service implements finance invoice producer gRPC server functions
type Service struct {
	metrics    stats.Metrics
	authorizer auth.AuthorizerInterface
}

// NewService initialises a new policy service
func NewService(
	metrics stats.Metrics,
	authorizer auth.AuthorizerInterface,
) *Service {
	return &Service{
		metrics:    metrics,
		authorizer: authorizer,
	}
}

func (s *Service) authorize(ctx context.Context, action auth.ActionOnResource) (bool, error) {
	return s.authorizer.IsAllowed(ctx, action)
}

// AuthorizationResponseObject Provides the definition for the graphql authorization response object
var AuthorizationResponseObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "authorization_response",
	Fields: graphql.Fields{
		"success": {
			Type:        graphql.Boolean,
			Description: "Indicates whether the request was successful",
		},
		"error_message": {
			Type:        graphql.String,
			Description: "If the request was unsuccessful, the response will include the error message",
		},
	},
})

// ToSchema provides the service as a graphql schema for go-graphql
func (s *Service) ToSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.getQuery(),
		Mutation: s.getMutation(),
	})
}

func (s *Service) getMutation() *graphql.Object {
	return nil
	// return graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "Mutation",
	// 	// Fields: graphql.Fields{},
	// })
}

func getResourceOwnerIdArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"resource_owner_id": {
			Type:         graphql.String,
			Description:  "The id of the owner of the resource. This field is optional.",
			DefaultValue: "",
		},
	}
}

func (s *Service) getQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"is_authorized_to_create_post": &graphql.Field{
				Type:        AuthorizationResponseObject,
				Name:        "is_authorized_to_run_tariff_assignment",
				Description: "Check if user has permissions to create the tariff assigments",
				Resolve:     s.IsAuthorizedToCreatePost,
			},
			"is_authorized_to_read_post": &graphql.Field{
				Name:        "is_authorized_to_run_tariff_assignment",
				Type:        AuthorizationResponseObject,
				Description: "Check if user has permissions to read the tariff assigments",
				Resolve:     s.IsAuthorizedToReadPost,
				Args:        getResourceOwnerIdArgs(),
			},
			"is_authorized_to_update_post": &graphql.Field{
				Type:        AuthorizationResponseObject,
				Name:        "is_authorized_to_run_tariff_assignment",
				Description: "Check if user has permissions to create the tariff assigments",
				Resolve:     s.IsAuthorizedToUpdatePost,
				Args:        getResourceOwnerIdArgs(),
			},
			"is_authorized_to_delete_post": &graphql.Field{
				Type:        AuthorizationResponseObject,
				Name:        "is_authorized_to_run_tariff_assignment",
				Description: "Check if user has permissions to create the tariff assigments",
				Resolve:     s.IsAuthorizedToDeletePost,
				Args:        getResourceOwnerIdArgs(),
			},
		},
	})
}
