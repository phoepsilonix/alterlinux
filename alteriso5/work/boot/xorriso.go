package boot

import (
	"log/slog"
	"os"
	"os/exec"
)

var Xorriso = xorriso{}

type xorriso struct {
	fsDir string
	out   string
	args  []*xorrisoArg
}

type xorrisoArg struct {
	bootMode Mode
	args     func(o *xorriso) []string
}

func (x *xorriso) addArgs(arg ...*xorrisoArg) {
	x.args = append(x.args, arg...)
}

func (x *xorriso) Args(bootmode ...Mode) *[]string {
	var args []string

	args = append(args, xorrisoCommonArgs.args(x)...)

	for _, arg := range x.args {
		if bootmode[0] == arg.bootMode {
			slog.Debug("Adding xorriso args", "mode", arg.bootMode, "args", arg.args)
			args = append(args, arg.args(x)...)
		}
	}
	return &args
}

func (x *xorriso) Build(dir string, out string, bootmode ...Mode) error {
	x.fsDir = dir
	x.out = out

	args := x.Args(bootmode...)

	cmd := exec.Command("xorriso", *args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	slog.Debug("Running xorriso", "args", cmd.Args)

	return cmd.Run()
}
