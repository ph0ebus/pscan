package plugins

import (
	"log/slog"
)

var logger = slog.Default()

func LogInit() *slog.Logger {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	return logger
}
