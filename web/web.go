package web

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/egustafson/websb-go/pkg/config"
	"github.com/egustafson/websb-go/web/api"
	"github.com/egustafson/websb-go/web/ui"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srvCfg := config.MustServerConfig(ctx)

	router := gin.Default() // root of the web routes
	router.GET("/live", LivenessHandler)
	router.GET("/ready", ReadinessHandler)
	router.GET("/healthz", ReadinessHandler)

	// TODO: middleware ??

	ui.Init(ctx, srvCfg, router.Group(""))      // initialize UI components
	api.Init(ctx, srvCfg, router.Group("/api")) // initialize API components

	// Start the server
	//
	var listener net.Listener
	{ // no-TLS
		slog.Info("starting server without tls")
		listener, err = net.Listen("tcp", fmt.Sprintf(":%d", srvCfg.Port))
		if err != nil {
			return fmt.Errorf("failed to start listener: %w", err)
		}
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", srvCfg.Port),
		Handler: router,
	}

	go func() {
		slog.Info("server listening", "port", srvCfg.Port)
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			cancel()
		}
	}()

	<-ctx.Done() // block waiting for shutdown signal
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer shutdownCancel()
	slog.Info("shutting down server")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown error", "error", err)
		srv.Close()
		return err
	}
	slog.Info("server shutdown complete")
	return nil
}
