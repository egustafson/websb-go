package config

import (
	"fmt"
	"log/slog"
	"time"
)

var (
	DefaultLogLevel = fmt.Sprintf("%d", slog.LevelInfo)
)

const (
	// DefaultPort is the default port number for the web server.
	DefaultPort = 8080

	// DefaultConfigFile is the default configuration file name.
	DefaultConfigFile = "websb.yaml"

	// ShutdownTimeout is the duration to wait for graceful shutdown.
	ShutdownTimeout = 5 * time.Second
)

// Registered environment variable names
const (
	// EnvLogLevel is the environment variable for setting log level.
	EnvLogLevel = "WEBSB_LOGLEVEL"
)
