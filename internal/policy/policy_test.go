package policy_test

import (
	"context"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/joepk90/graphql-auth/internal/policy"
	"github.com/stretchr/testify/assert"
)

func TestIsAuthorizedTCreatePost_Successfully(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_get_tariff_assignment_progress"

	ts.metricsMock.EXPECT().ObserveQuery(queryType, method)
	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.GetTariffAssignmentProcess("*"))

	resp, err := ts.svc.IsAuthorizedToCreatePost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedTCreatePost_AuthorizationError(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_get_tariff_assignment_progress"
	errorMsg := "Principal not authorized to get tariff assignment progress"

	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.GetTariffAssignmentProgress("*")).Return("", errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, method)

	resp, err := ts.svc.IsAuthorizedToCreatePost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedTReadPost_Successfully(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_get_tariff_assignment_progress"

	ts.metricsMock.EXPECT().ObserveQuery(queryType, method)
	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.GetTariffAssignmentProcess("*"))

	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedTReadPost_AuthorizationError(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_get_tariff_assignment_progress"
	errorMsg := "Principal not authorized to get tariff assignment progress"

	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.GetTariffAssignmentProgress("*")).Return("", errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, method)

	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedToUpdatePost_Successfully(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_run_tariff_assignment_process"

	ts.metricsMock.EXPECT().ObserveQuery(queryType, method)
	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.RunTariffAssignmentProcess())

	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedToUpdatePost_AuthorizationError(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_run_tariff_assignment_process"
	errorMsg := "Principal not authorized to run tariff assignments process"

	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.RunTariffAssignmentProcess()).Return("", errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, method)

	resp, err := ts.svc.IsAuthorizedToReadPost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}

func TestIsAuthorizedToDeletePost_Successfully(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_run_tariff_assignment_process"

	ts.metricsMock.EXPECT().ObserveQuery(queryType, method)
	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.RunTariffAssignmentProcess())

	resp, err := ts.svc.IsAuthorizedToDeletePost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.NoError(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(nil))
}

func TestIsAuthorizedToDeletePost_AuthorizationError(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	method := "is_authorized_to_run_tariff_assignment_process"
	errorMsg := "Principal not authorized to run tariff assignments process"

	// request policy decision client
	// ts.authorizerMock.EXPECT().IsAllowed(gomock.Any(), auth.RunTariffAssignmentProcess()).Return("", errors.New("not allowed"))
	ts.metricsMock.EXPECT().ObserveQueryError(queryType, method)

	resp, err := ts.svc.IsAuthorizedToDeletePost(graphql.ResolveParams{
		Context: ctx,
	})

	assert.Error(t, err)
	assert.Equal(t, resp, policy.MapAuthenticationResponse(&errorMsg))
}
