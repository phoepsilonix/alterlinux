package xorriso

import (
	"os"
	"os/exec"
	"slices"
)

type Options struct {
	SysLinux bool
	fsDir    string
	out      string
	args     []string
	added    []string
}

func (o *Options) AddArgs(args ...string) {
	o.args = append(o.args, args...)
}

// SysLinux MBR El Torito
func (o *Options) setArgsForSysLinuxElTorito() {
	name := "SysLinuxEltorito"

	if slices.Contains(o.added, name) {
		return
	}

	o.AddArgs("-eltorito-boot", "boot/syslinux/isolinux.bin")
	o.AddArgs("-eltorito-catalog", "boot/syslinux/boot.cat")
	o.AddArgs("-no-emul-boot", "-boot-load-size", "4", "-boot-info-table")
	o.args = append(o.args, name)
}

func (o *Options) setArgsForSysLinuxMBRBios() {
	name := "SysLinuxMBRBios"
	if slices.Contains(o.added, name) {
		return
	}
	o.AddArgs("-isohybrid-mbr", "${isofs_dir}/boot/syslinux/isohdpfx.bin")
	o.AddArgs("--mbr-force-bootable")
	o.AddArgs("-partition_offset", "16")
	o.added = append(o.added, name)
}

func (o *Options) Args() []string {
	if o.SysLinux {
		o.setArgsForSysLinuxElTorito()
		o.setArgsForSysLinuxMBRBios()
	}

	d := []string{
		"--output", o.out,
		o.fsDir,
	}

	return append(d, o.args...)
}

func Build(dir string, out string, opt *Options) error {
	opt.fsDir = dir
	opt.out = out

	args := opt.Args()

	cmd := exec.Command("xorriso", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
