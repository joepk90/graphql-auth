package policy_test

import (
	"testing"

	"github.com/joepk90/graphql-auth/internal/mocks"
	"github.com/joepk90/graphql-auth/internal/policy"
	"go.uber.org/mock/gomock"
)

const (
	queryType = "query"
	// mutationQueryType = "mutation"
)

type testSuite struct {
	svc            *policy.Service
	metricsMock    *mocks.MockMetrics
	authorizerMock *mocks.MockAuthorizerInterface
}

func testSetup(t *testing.T) (testSuite, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	ts := testSuite{
		metricsMock:    mocks.NewMockMetrics(ctrl),
		authorizerMock: mocks.NewMockAuthorizerInterface(ctrl),
	}

	ts.svc = policy.NewService(ts.metricsMock, ts.authorizerMock)

	return ts, ctrl
}
