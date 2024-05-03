package build

import (
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work"
)

func build() error {
	current, err := os.Getwd()
	if err != nil {
		return err
	}

	workDir := path.Join(current, "work")
	outDir := path.Join(current, "out")

	work, err := work.New(workDir)
	if err != nil {
		return err
	}

	// Dummy profile
	profile := config.Profile{
		Base:       path.Join(current, "profile"),
		InstallDir: "alter",
		BootModes:  []string{"SysLinux"},
	}

	// TODO: Add more targets
	target := config.NewTarget("x86_64", outDir)
	return work.Build(profile, target)
}
