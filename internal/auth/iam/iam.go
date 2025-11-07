package iam

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	incomingCtxKey = contextKey("user.incoming")
	outgoingCtxKey = contextKey("user.outgoing")
)

const (
	header       = "authorization"
	headerScheme = "bearer"
)

// Model represents the principal in the authorization setting.
// The princiapl is currently only identified by a JWT issued
type Model struct {
	Token string
}

// New initialises a new principal.
func New(token string) *Model {
	return &Model{
		Token: token,
	}
}

// FromIncomingCtx retrieves an incoming user principal.
func FromIncomingCtx(ctx context.Context) string {
	if tkn := ctx.Value(incomingCtxKey); tkn != nil {
		return tkn.(string)
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if val := md.Get(header); len(val) > 0 {
			if splits := strings.SplitN(val[0], " ", 2); len(splits) == 2 {
				return splits[1]
			}
		}
	}
	return ""
}

func PrincipalFromCtx(ctx context.Context) *Model {
	if pr := FromIncomingCtx(ctx); pr != "" {
		return &Model{Token: pr}
	}
	if pr := FromIncomingCtx(ctx); pr != "" {
		return &Model{Token: pr}
	}
	return &Model{}
}

func ToIncomingCtx(ctx context.Context, tkn string) context.Context {
	return context.WithValue(ctx, incomingCtxKey, tkn)
}

func ToOutgoingCtx(ctx context.Context, tkn string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, header, fmt.Sprintf("%s %s", headerScheme, tkn))
	return context.WithValue(ctx, outgoingCtxKey, tkn)
}
