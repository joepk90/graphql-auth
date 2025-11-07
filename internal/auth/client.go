package auth

import (
	"context"
	"fmt"

	httpClient "github.com/joepk90/graphql-auth/internal/http"
)

const (
	path       = "/resource/"
	methodPost = "POST"
)

type AuthService struct {
	client httpClient.HttpService
}

func NewAuthService(authURL string) *AuthService {
	return &AuthService{
		client: *httpClient.NewHttpService(authURL),
	}
}

type AuthorizeActionsResponse struct {
	IsAuthorized bool   `json:"isAuthorized,omitempty"`
	Action       string `json:"action"`
}

func (a AuthService) NewPostRequestWithContext(ctx context.Context, reqBody any) (*AuthorizeActionsResponse, error) {

	var resBody []AuthorizeActionsResponse
	err := a.client.NewRequestWithContext(ctx, methodPost, path, reqBody, &resBody)

	if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}

	// this is a slight hack
	return &resBody[0], nil

}
