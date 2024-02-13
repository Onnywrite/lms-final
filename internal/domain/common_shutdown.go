package domain

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Stoppable interface {
	Stop(context.Context)
}

func ShutdownOnSignal(timeout time.Duration, log *slog.Logger, processes ...Stoppable) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	log.Debug("Waiting for shutdown")

	<-shutdown
	log.Info("Shutting down")

	wg := sync.WaitGroup{}
	for _, p := range processes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			p.Stop(ctx)
			select {
			case <-ctx.Done():
				log.Error("Shutdown timeout exceeded")
			}
		}()
	}

	wg.Wait()
	log.Info("Gracefully stopped")
}
