package context

import (
	"context"
	"time"
)

const (
	requestedAtKey        = contextKey("requested-at")
)

type contextKey string

func (ckt contextKey) String() string {
	return string(ckt)
}

func WithRequestedAt(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, requestedAtKey, t)
}
