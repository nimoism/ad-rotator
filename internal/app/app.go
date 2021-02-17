package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	intgrpc "github.com/nimoism/ad-rotator/internal/api/grpc"
	"github.com/nimoism/ad-rotator/internal/log"
	intservice "github.com/nimoism/ad-rotator/internal/service"
)

type App struct {
	config Config
	log    log.Logger
}

func NewApp(config Config) (*App, error) {
	logger, err := NewLogger(config.Logger)
	if err != nil {
		return nil, fmt.Errorf("logger initialization error: %w", err)
	}
	logger.Debug("logger initialized")

	app := &App{
		config: config,
		log:    logger,
	}
	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	conf := a.config
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	repo, err := createRepo(ctx, a.log, conf.Storage)
	if err != nil {
		return fmt.Errorf("app starting error: %w", err)
	}
	a.log.Debug("Repository initialized")

	service := intservice.NewService(a.log, repo)
	a.log.Debug("Services initialized")

	api := intgrpc.NewAPIServer(a.log, service)
	a.log.Debug("GRPC service initialized")

	listener, err := net.Listen("tcp", conf.API.Listen)
	if err != nil {
		return fmt.Errorf("app bind network error: %w", err)
	}

	go func() {
		signals := make(chan os.Signal, syscall.SIGQUIT|syscall.SIGTERM)
		signal.Notify(signals)
		select {
		case s := <-signals:
			a.log.Info("Got signal %v", s)
			cancel()
		case <-ctx.Done():
		}
		signal.Stop(signals)
		api.Stop()
	}()

	a.log.Debug("Running server...")
	return api.Serve(listener)
}
