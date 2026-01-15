package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/egustafson/websb-go/pkg/config"
	"github.com/egustafson/websb-go/web"
)

func Start(flags config.Flags) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// hook signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		slog.Warn("received signal, shutting down", "signal", sig.String())
		cancel()
	}()

	_, ctx, err := config.InitServerConfig(ctx, flags)
	if err != nil {
		slog.Error("error initializing configuration", "error", err)
		return err
	}

	// TODO: initialize other subsystems here

	return web.Run(ctx) // blocks until shutdown
}
