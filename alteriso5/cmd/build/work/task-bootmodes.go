package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
)

var makeSysLinux = NewBuildTask("makeSysLinux", func(w *Work) error {

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

	if err := utils.CopyAll(cpFiles...); err != nil {
		return err
	}

	return nil
})

var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w *Work) error {

	return w.RunOnce(makeSysLinux)
})
