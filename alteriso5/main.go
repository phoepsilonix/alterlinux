package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/work"
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
	target := config.NewTarget([]string{"x86_64"}, outDir)
	return work.Build(profile, target)
}

func main() {
	if err := check(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := build(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
