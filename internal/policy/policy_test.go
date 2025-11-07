package policy_test

import (
	"context"
	"errors"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/joepk90/graphql-auth/internal/auth"
	"github.com/joepk90/graphql-auth/internal/policy"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIsAuthorizedTCreatePost_Successfully(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	// request policy decision client
	ts.metricsMock.EXPECT().ObserveQuery(queryType, policy.CreatePostRequest)
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanCreatePost()).Return(true, nil)

	// act
	resp, err := ts.svc.IsAuthorizedToCreatePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedTCreatePost_AuthorizationError(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	errorMsg := policy.CreatePostErrorMsg

	// request policy decision client
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanCreatePost()).Return(false, errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, policy.CreatePostRequest)

	// act
	resp, err := ts.svc.IsAuthorizedToCreatePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedTReadPost_Successfully(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	// request policy decision client
	ts.metricsMock.EXPECT().ObserveQuery(queryType, policy.ReadPostRequest)
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanReadPost("*"))

	// act
	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedTReadPost_AuthorizationError(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	errorMsg := policy.ReadPostErrorMsg

	// request policy decision client
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanReadPost("*")).Return(false, errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, policy.ReadPostRequest)

	// act
	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedToUpdatePost_Successfully(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	// request policy decision client
	ts.metricsMock.EXPECT().ObserveQuery(queryType, policy.UpdatePostRequest)
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanUpdatePost("*"))

	// act
	resp, err := ts.svc.IsAuthorizedToUpdatePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedToUpdatePost_AuthorizationError(t *testing.T) {
	// arrange
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	errorMsg := policy.UpdatePostErrorMsg

	// request policy decision client
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanUpdatePost("*")).Return(false, errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, policy.UpdatePostRequest)

	// act
	resp, err := ts.svc.IsAuthorizedToUpdatePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedToDeletePost_Successfully(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	// request policy decision client
	ts.metricsMock.EXPECT().ObserveQuery(queryType, policy.DeletePostRequest)
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanDeletePost("*"))

	// act
	resp, err := ts.svc.IsAuthorizedToDeletePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedToDeletePost_AuthorizationError(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	errorMsg := policy.DeletePostErrorMsg

	// request policy decision client
	ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.CanDeletePost("*")).Return(false, errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, policy.DeletePostRequest)

	// act
	resp, err := ts.svc.IsAuthorizedToDeletePost(graphql.ResolveParams{
		Context: ctx,
	})

	// assert
	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}
