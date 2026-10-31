package app

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/Skoowshot/qrix-runtime/internal/app/server"
)

type App struct {
	servers []Server
}

func NewApp() *App {
	return &App{
		servers: make([]Server, 0),
	}
}

func (app *App) Start() error {
	ctrlServer, err := server.ListenControl(":7178")
	if err != nil {
		return fmt.Errorf("create control server: %w", err)
	}

	app.servers = append(
		app.servers,
		ctrlServer,
	)

	for _, s := range app.servers {
		go func() {
			if err := s.Start(); err != nil {
				slog.Error("failed to start server", "err", err)
			}
		}()
	}

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	stopped := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, s := range app.servers {
			wg.Add(1)

			go func(srv Server) {
				defer wg.Done()
				srv.GracefulStop()
			}(s)
		}
		wg.Wait()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		for _, s := range app.servers {
			s.Stop()
		}
		return ctx.Err()
	case <-stopped:
		return nil
	}
}
