package logger

import (
	"context"
	"io"
	"log/slog"
)

type TraceIdFn func(ctx context.Context)

type Logger struct {
	log       slog.Handler
	traceIdFn TraceIdFn
}

func New(w io.Writer, minLevel Level, servieName string, traceIdFn TraceIdFn) *Logger {
	return new(w, minLevel, serviceName, traceIdFn, Events{})
}

func new(w io.Writer, Level )
