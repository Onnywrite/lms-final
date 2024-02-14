package domain

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

const (
	LocalEnv = "loc"
	DevEnv   = "dev"
	ProdEnv  = "prod"
)

func MustSetupLoggerInDir(env, logsDir string) *slog.Logger {
	return SetupLogger(env, MustCreatelogFile(logsDir))
}

func SetupLogger(env string, out *os.File) *slog.Logger {
	switch env {
	case LocalEnv:
		return slog.New(
			slog.NewTextHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DevEnv:
		return slog.New(
			slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ProdEnv:
		return slog.New(
			slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic("unexpected env value")
	}
}

func MustCreatelogFile(logsDir string) (f *os.File) {
	err := os.MkdirAll(logsDir, 0666)
	if err != nil {
		panic(err)
	}

	logFileName := fmt.Sprintf("%s.log", time.Now().Format("15-04-05_01-02-2006"))
	f, err = os.Create(logsDir + "/" + logFileName)
	if err != nil {
		panic(err)
	}
	return
}
