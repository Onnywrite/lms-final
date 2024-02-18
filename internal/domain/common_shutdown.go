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
	Stop(context.Context, chan<- error)
}

func ShutdownOnSignal(timeout time.Duration, logger *slog.Logger, processes ...Stoppable) {
	const op = "domain.ShutdownOnSignal"
	log := logger.With(slog.String("op", op))

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
			done := make(chan error)
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			p.Stop(ctx, done)
			select {
			case err := <-done:
				if err != nil {
					log.Error("Shutdown exited with error", slog.String("err", err.Error()))
				}
			case <-ctx.Done():
				log.Error("Shutdown timeout exceeded")
			}
		}()
	}

	wg.Wait()
	log.Info("Gracefully stopped")
}
