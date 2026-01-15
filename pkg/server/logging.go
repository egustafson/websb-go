package server

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/egustafson/websb-go/pkg/config"
)

func init() { // bootstrap logging
	logWr := os.Stderr
	levelStr := config.DefaultLogLevel

	logger := slog.New(slog.NewTextHandler(logWr, &slog.HandlerOptions{
		Level: strToLevel(levelStr),
	}))
	slog.SetDefault(logger)
	slog.Debug("logging initialized", "level", levelStr)
}

func strToLevel(levelStr string) slog.Level {
	if i, err := strconv.Atoi(levelStr); err == nil {
		return slog.Level(i)
	}
	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
