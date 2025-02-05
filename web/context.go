package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	writerKey ctxKey = iota + 1
	traceIDKey
)

func setTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}
