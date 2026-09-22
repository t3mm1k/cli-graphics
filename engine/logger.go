package engine

import (
	"io"
	"log/slog"
	"os"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

var Log Logger = slog.New(slog.NewTextHandler(io.Discard, nil))

func enableFileLogging(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	Log = slog.New(handler)
	return file, nil
}

func resetLogger() {
	Log = slog.New(slog.NewTextHandler(io.Discard, nil))
}

func SetLogger(l Logger) {
	if l != nil {
		Log = l
	}
}
