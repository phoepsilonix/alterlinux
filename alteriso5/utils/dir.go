package utils

import (
	"log/slog"
	"os"
)

func MkdirsAll(dirs ...string) error {
	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		slog.Debug("Creating directory", "dir", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			slog.Error("Failed to create directory", "dir", dir, "error", err)
			return err
		}
	}
	return nil
}
