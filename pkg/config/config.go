package config

import (
	"context"
	"crypto/sha1"
	"fmt"
	"reflect"
)

// ServerConfig holds configuration settings for the web server.
type ServerConfig struct {

	// Flags contains command-line flag values.
	Flags Flags `yaml:"-" json:"-"`

	// Port is the port number the server listens on.
	Port int `yaml:"port" json:"port"`
}

// Flags holds command-line flag values.
type Flags struct {
	Verbose bool
}

// InitServerConfig initializes the ServerConfig from the given flags
// and returns it along with an updated context containing the config.
func InitServerConfig(ctx context.Context, flags Flags) (*ServerConfig, context.Context, error) {
	cfg := &ServerConfig{ // default config
		Port: DefaultPort,
	}
	cfg.Flags = flags // apply flags

	if ctx != nil {
		ctx = setServerConfig(ctx, cfg)
	}
	return cfg, ctx, nil
}

var (
	// serverConfigTypeID is a unique identifier for ServerConfig type.
	serverConfigTypeID = sha1.Sum(
		[]byte(fmt.Sprintf("%s.%s",
			reflect.TypeOf(ServerConfig{}).PkgPath(),
			reflect.TypeOf(ServerConfig{}).Name(),
		)),
	)
)

// MustServerConfig retrieves the ServerConfig from the context.
// It panics if the config is not found.
func MustServerConfig(ctx context.Context) *ServerConfig {
	cfg, ok := ctx.Value(serverConfigTypeID).(*ServerConfig)
	if !ok || cfg == nil {
		panic("server config not found in context")
	}
	return cfg
}

// setServerConfig sets the ServerConfig in the context.
func setServerConfig(ctx context.Context, cfg *ServerConfig) context.Context {
	return context.WithValue(ctx, serverConfigTypeID, cfg)
}
