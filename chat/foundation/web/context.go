package web

import (
	"context"
	"net/http"

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

func GetTraceID(ctx context.Context) uuid.UUID {
	if v, ok := ctx.Value(traceIDKey).(uuid.UUID); !ok {
		return uuid.UUID{}
	} else {
		return v
	}
}

func setWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, traceIDKey, w)
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	v, ok := ctx.Value(ctx).(http.ResponseWriter)
	if !ok {
		return nil
	}

	return v
}
