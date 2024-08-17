package utils

import (
	"log/slog"
	"os/exec"

	"github.com/Hayao0819/nahi/exutils"
)

func CommandWithStdio(cmd string, args ...string) (*exec.Cmd) {
	slog.Debug("CommandWithStdio", "cmd", cmd, "args", args)
	return exutils.CommandWithStdio(cmd, args...)
}
