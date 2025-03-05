package logger

import (
	"log/slog"
	"time"
)

type Level slog.Level

const (
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
)

type Record struct {
	Time    time.Time
	Message string
	Level   Level
	Attr    map[string]any
}
