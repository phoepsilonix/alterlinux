package airootfs

import (
	"os"
	"os/exec"
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
		if os.IsNotExist(err) {
			env.initilized = false
			return &env, nil
		} else {
			return nil, err
		}
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

	pacstrap := exec.Command("pacstrap", "-c", e.Dir, "base", "base-devel", "linux", "linux-firmware", "syslinux", "mkinitcpio-archiso")
	pacstrap.Env = append(os.Environ(), "LANG=C")
	pacstrap.Stdout = os.Stdout
	pacstrap.Stderr = os.Stderr
	if err := pacstrap.Run(); err != nil {
		return err
	}

	e.initilized = true

	return nil
}

type kernel struct {
	Linux  string
	Initrd string
}

// func (e *Chroot) FindKernels() ([]kernel, error) {
// 	kernels := []kernel{}

// 	presetsDir := path.Join(e.Dir, "etc", "mkinitcpio.d")
// 	entry, err := os.ReadDir(presetsDir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, e := range entry {
// 		if e.IsDir() || !strings.HasSuffix(e.Name(), ".preset") {
// 			continue
// 		}

// 		fp := path.Join(presetsDir, e.Name())
// 		env, err := utils.LoadEnvFile(fp)
// 		if err != nil {
// 			continue
// 		}

// 		ker := env["ALL_kver"]
// 		initrd := env["default_image"]

// 		if ker != "" && initrd != "" {
// 			kernels = append(kernels, kernel{
// 				Linux:  ker,
// 				Initrd: initrd,
// 			})
// 		}

// 	}

// 	slog.Debug("FindKernels:", "kernels", kernels)
// 	return kernels, nil
// }

func (e *Chroot) FindKernels() ([]kernel, error) {
	return []kernel{
		{
			Linux:  "/boot/vmlinuz-linux",
			Initrd: "/boot/initramfs-linux.img",
		},
	}, nil
}
