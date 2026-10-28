package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Skoowshot/qrix-runtime/internal/app"
)

const stopTimeout = 5 * time.Second

func main() {
	a := app.NewApp()

	if err := a.Start(); err != nil {
		slog.Error("failed to start app", "err", err)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()

	if err := a.Stop(ctx); err != nil {
		slog.Error("failed to stop app", "err", err)
	}

}
