package log

import (
	"log/slog"

	"github.com/m-mizutani/clog"
)

var logger = slog.New(clog.New(
	clog.WithLevel(slog.LevelDebug),
	clog.WithColor(true),
))

func Setup() {
	slog.SetDefault(logger)
}
