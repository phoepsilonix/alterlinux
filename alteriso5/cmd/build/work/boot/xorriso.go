package boot

import (
	"os"
	"os/exec"
	"path"
)

var Xorriso = xorriso{}

type xorriso struct {
	fsDir string
	out   string
	args  []*xorrisoArg
}

type xorrisoArg struct {
	name string
	args *[]string
}

func (xa *xorrisoArg) add(args ...string) {
	if xa.args == nil {
		xa.args = &[]string{}
	}

	arg := append(*(xa.args), args...)
	xa.args = &arg
}

func (x *xorriso) hasArg(a *xorrisoArg) bool {
	for _, x := range x.args {
		if x.name == a.name {
			return true
		}
	}
	return false
}

func (o *xorriso) addArg(arg *xorrisoArg) {
	if !o.hasArg(arg) {
		o.args = append(o.args, arg)
	}
}

func (o *xorriso) defaultArgs() *xorrisoArg {

	out := path.Join(o.out, "alterlinux.iso")

	d := []string{
		"-as", "mkisofs",
		"-iso-level", "3",
		"-full-iso9660-filenames",
		"-joliet",
		"-joliet-long",
		"-rational-rock",
		"--output", out,
		o.fsDir,
	}
	return &xorrisoArg{
		name: "default",
		args: &d,
	}
}

func (x *xorriso) Args() *[]string {
	args := []string{}
	for _, a := range x.args {
		args = append(args, *a.args...)
	}

	def := x.defaultArgs()
	args = append(args, *def.args...)
	return &args
}

func (x *xorriso) Build(dir string, out string) error {
	x.fsDir = dir
	x.out = out

	args := x.Args()

	cmd := exec.Command("xorriso", *args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
