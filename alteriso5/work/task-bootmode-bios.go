package work

import (
	"log/slog"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/FascodeNet/alterlinux/alteriso5/work/boot"
)

// SysLinux
var makeBiosSysLinuxMbr = NewBuildTask("makeBiosSysLinuxMbr", func(w Work) error {
	slog.Debug("Setting up SYSLINUX for BIOS booting from a disk...")

	// Get directories
	dirs := w.Dirs
	isoSyslinuxDir := path.Join(dirs.Iso, "boot", "syslinux")
	biosFilesDir := path.Join(dirs.Pacstrap, "usr", "lib", "syslinux", "bios")

	// Create directories
	if err := utils.MkdirsAll(isoSyslinuxDir, dirs.SyslinuxConfig); err != nil {
		return err
	}

	// syslinux config
	orgSyslinuxConfigDir := ""
	if w.profile.UseAlterSysLinux {
		orgSyslinuxConfigDir = path.Join(dirs.Data, "syslinux")
	} else {
		orgSyslinuxConfigDir = path.Join(w.profile.Base, "syslinux", "bios")
	}
	sc, err := boot.ReadSysLinuxConf(orgSyslinuxConfigDir)
	if err != nil {
		return err
	}
	sysLinuxConfigDir := path.Join(dirs.Work, w.target.Arch, "syslinux")
	if err := os.MkdirAll(sysLinuxConfigDir, 0755); err != nil {
		return err
	}
	if err := sc.ParseAndBuild(w.Values(), sysLinuxConfigDir); err != nil {
		return err
	}

	// Copy files
	cpFiles := []utils.CopyTask{
		{
			Source: biosFilesDir,
			Dest:   isoSyslinuxDir,
			Skip:   utils.OnlySpecificExtention(".c32"),
			Perm:   0644,
		},
		{
			Source: dirs.SyslinuxConfig,
			Dest:   isoSyslinuxDir,
		},
		{
			Source: path.Join(biosFilesDir, "lpxelinux.0"),
			Dest:   path.Join(isoSyslinuxDir, "lpxelinux.0"),
		},
		{
			Source: path.Join(biosFilesDir, "memdisk"),
			Dest:   path.Join(isoSyslinuxDir, "memdisk"),
		},
	}

	chroot, err := w.GetChroot()
	if err != nil {
		return err
	}
	kernels, err := chroot.FindKernels()
	if err != nil {
		return err
	}

	for _, k := range kernels {
		cpFiles = append(cpFiles, utils.CopyTask{
			Source: path.Join(w.Dirs.Pacstrap, k.Linux),
			Dest:   path.Join(dirs.Iso, "boot", w.target.Arch, path.Base(k.Linux)),
			Perm:   0644,
		}, utils.CopyTask{
			Source: path.Join(w.Dirs.Pacstrap, k.Initrd),
			Dest:   path.Join(dirs.Iso, "boot", w.target.Arch, path.Base(k.Initrd)),
			Perm:   0644,
		})
	}

	if err := utils.CopyAll(cpFiles...); err != nil {
		return err
	}

	return nil
})

var makeBiosSysLinuxElTorito = NewBuildTask("makeBiosSysLinuxElTorito", func(w Work) error {
	return nil
})
