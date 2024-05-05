package airootfs

import (
	"os"
	"os/exec"
	"path/filepath"
)

type SquashFS struct {
	Base string
	Out  string
	Args []string
}

func (s *SquashFS) Build() error {
	args := append([]string{s.Base, s.Out}, s.Args...)

	if err := os.MkdirAll(filepath.Dir(s.Out), 0755); err != nil {
		return err
	}

	cmd := exec.Command("mksquashfs", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
