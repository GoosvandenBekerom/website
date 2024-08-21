package logger

import (
	"log/slog"
)

func Initialize() {
	// TODO: save warn/error logs to file/db as well
	slog.SetDefault(slog.New(NewPrettyPrintHandler(slog.LevelInfo)))
}
