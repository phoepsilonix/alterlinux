package airootfs

import (
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
)

type Chroot struct {
	Arch       string
	Dir        string
	initilized bool
}

func GetChrootDir(dir, arch string) (*Chroot, error) {
	env := Chroot{
		Arch: arch,
		Dir:  dir,
	}

	entry, err := os.ReadDir(dir)
	if err != nil {
		return nil, err

	}

	if len(entry) > 0 {
		env.initilized = true
	}

	return &env, nil
}

func (e *Chroot) Init() error {
	if err := os.MkdirAll(e.Dir, 0755); err != nil {
		return err
	}
	if e.initilized {
		return nil
	}

	pacstrap := exec.Command("pacstrap", "-c", e.Dir, "base", "base-devel", "linux", "linux-firmware", "syslinux")
	pacstrap.Env = append(os.Environ(), "LANG=C")
	pacstrap.Stdout = os.Stdout
	pacstrap.Stderr = os.Stderr
	if err := pacstrap.Run(); err != nil {
		return err
	}

	e.initilized = true

	return nil
}

func (e *Chroot) FindKernels() ([]string, error) {
	kernels := []string{}

	//bootDir := path.Join(e.Dir, "boot")

	presetsDir := path.Join(e.Dir, "etc", "mkinitcpio.d")
	entry, err := os.ReadDir(presetsDir)
	if err != nil {
		return nil, err
	}

	for _, e := range entry {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".preset") {
			continue
		}

		fp := path.Join(presetsDir, e.Name())
		env, err := utils.LoadEnvFile(fp)
		if err != nil {
			continue
		}

		kernel := env["ALL_kver"]

		if kernel != "" {
			kernels = append(kernels, kernel)
		}

	}

	slog.Debug("FindKernels:", "kernels", kernels)
	return kernels, nil
}
