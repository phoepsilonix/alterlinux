package log

import (
	"log/slog"
	"os"

	"github.com/m-mizutani/clog"
)

var logger = slog.New(clog.New(
	clog.WithLevel(slog.LevelDebug),
	clog.WithColor(true),
	clog.WithWriter(os.Stderr),
))

func Setup() {
	slog.SetDefault(logger)
}
