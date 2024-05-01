package build

import (
	"errors"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work"
)

func check() error {
	if os.Getuid() != 0 {
		return errors.New("this program must be run as root")
	}
	return nil
}

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

	profile := config.DummyProfile()
	target := config.NewTarget("x86_64", outDir)
	return work.Build(profile, target)
}
