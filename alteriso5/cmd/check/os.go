package check

import (
	"errors"
	"log/slog"
	"runtime"

	"github.com/Hayao0819/go-distro"
)

var ErrUnsupportedOS = errors.New("unsupported OS")

func OS() error {
	if runtime.GOOS != "linux" {
		return ErrUnsupportedOS
	}

	osdetail := distro.GetDetail()
	if id := osdetail.ID(); id != "arch" {
		slog.Warn("Unsupported OS: " + string(id))
	}

	return nil
}
