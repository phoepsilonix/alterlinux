package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
)

var makeSysLinux = NewBuildTask("makeSysLinux", func(w Work) error {

	slog.Debug("Setting up SYSLINUX for BIOS booting from a disk...")
	dirs := w.GetDirs()

	isoSyslinuxDir := path.Join(dirs.Iso, "boot", "syslinux")

	if err := utils.MkdirsAll(isoSyslinuxDir, dirs.SyslinuxConfig); err != nil {
		return err
	}

	biosFilesDir := path.Join(dirs.Pacstrap, "usr", "lib", "syslinux", "bios")
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
			Source: path.Join(w.GetDirs().Pacstrap, k.Linux),
			Dest:   path.Join(dirs.Iso, "boot", path.Base(k.Linux)),
			Perm:   0644,
		}, utils.CopyTask{
			Source: path.Join(w.GetDirs().Pacstrap, k.Initrd),
			Dest:   path.Join(dirs.Iso, "boot", path.Base(k.Initrd)),
			Perm:   0644,
		})
	}

	if err := utils.CopyAll(cpFiles...); err != nil {
		return err
	}

	return nil
})

var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w Work) error {
	return w.RunOnce(makeSysLinux)
})
