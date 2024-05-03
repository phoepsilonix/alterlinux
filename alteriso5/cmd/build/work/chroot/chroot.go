package chroot

import (
	"os"
	"os/exec"
)

type Env struct {
	Arch string
	Dir  string
}

func New(dir, arch string) *Env {
	return &Env{
		Arch: arch,
		Dir:  dir,
	}
}

func (e *Env) Init() error {
	if err := os.MkdirAll(e.Dir, 0755); err != nil {
		return err
	}

	pacstrap := exec.Command("pacstrap", "-c", e.Dir, "base", "base-devel", "linux", "linux-firmware", "syslinux")
	pacstrap.Env = append(os.Environ(), "LANG=C")
	pacstrap.Stdout = os.Stdout
	pacstrap.Stderr = os.Stderr
	if err := pacstrap.Run(); err != nil {
		return err
	}

	return nil
}
