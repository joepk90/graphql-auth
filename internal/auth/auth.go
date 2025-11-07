package auth

import (
	"context"
	"fmt"
)

type AuthorizerInterface interface {
	IsAllowed(context.Context, ActionOnResource) (bool, error)
}

type AuthorizationError struct {
	action         string
	principalEmail string
	resource       string
}

func (e AuthorizationError) Error() string {
	return fmt.Sprintf("Principal %s not allowed to perform action %s on resource %s",
		e.principalEmail, e.action, e.resource)
}

func NewAuthorizationError(action, resource string) AuthorizationError {
	return AuthorizationError{
		action:   action,
		resource: resource,
	}
}

type Authorizer struct {
	authClient *AuthService
}

func NewAuthorizer(authClient *AuthService) Authorizer {
	return Authorizer{
		authClient: authClient,
	}
}

// ResourceIdentity encapsulates the identity of a resource - the kind (type) and a unique identifier.
type ResourceIdentity struct {
	Kind string
	ID   string
}

// Resource represents a single identifiable resource, and any relevant attributes.
type Resource struct {
	ResourceIdentity
	Attr map[string]any
}

// NewResource initialises a new resource.
func NewResource(kind, id string) *Resource {
	return &Resource{
		ResourceIdentity: ResourceIdentity{
			Kind: kind,
			ID:   id,
		},
	}
}

type AuthorizeRequest struct {
	ResourceType    string   `json:"resourceType"`
	ActionTypes     []string `json:"actionTypes"`
	ResourceID      string   `json:"resourceId"`
	ResourceOwnerId string   `json:"resourceOwnerId"`
}

func newAuthizationRequest(action ActionOnResource) AuthorizeRequest {
	return AuthorizeRequest{
		ResourceType:    action.Resource,
		ActionTypes:     []string{action.Action},
		ResourceID:      action.ID,
		ResourceOwnerId: "*",
	}
}

func (a *Authorizer) IsAllowed(ctx context.Context, action ActionOnResource) (bool, error) {
	request := newAuthizationRequest(action)

	response, err := a.authClient.NewPostRequestWithContext(ctx, request)

	if err != nil {
		return false, fmt.Errorf("failed to check if principal is allowed to %s on resource: %s:%w",
			action.Action, action.Resource, err)
	}

	if !response.IsAuthorized {
		return false, NewAuthorizationError(
			action.Action,
			action.Resource,
		)
	}

	return true, nil
}
