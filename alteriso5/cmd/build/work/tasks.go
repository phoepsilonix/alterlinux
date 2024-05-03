package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/airootfs"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/boot"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/chroot"
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
)

var makeBaseDirs *BuildTask = NewBuildTask("makeBaseDirs",
	func(work *Work) error {

		dirs := []string{
			work.Base,
			work.target.Out,
		}

		if err := utils.MkdirsAll(dirs...); err != nil {
			return err
		}
		return nil
	})

var makeChroot *BuildTask = NewBuildTask("makeChroot", func(work *Work) error {

	dir := path.Join(work.Base, work.target.Arch, "airootfs")
	env := chroot.New(dir, work.target.Arch)
	if err := env.Init(); err != nil {
		return err

	}
	return nil
})

var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w *Work) error {

	makeSysLinux := NewBuildTask("makeSysLinux", func(w *Work) error {

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

	return makeSysLinux.Run(w)
})

var makeAirootfs *BuildTask = NewBuildTask("makeAirootfs", func(w *Work) error {

	slog.Debug("Copying profile to airootfs...")
	airootfsDir := path.Join(w.Base, w.target.Arch, "airootfs")
	isoDir := path.Join(w.Base, "iso")

	sqfs := airootfs.SquashFS{
		Base: airootfsDir,
		Out:  path.Join(isoDir, w.profile.InstallDir, w.target.Arch, "airootfs.sfs"),
	}

	return sqfs.Build()
})

var makeOutFiles *BuildTask = NewBuildTask("makeOutFiles", func(w *Work) error {

	boot.Xorriso.SetArgsForSysLinuxElTorito()
	boot.Xorriso.SetArgsForSysLinuxElTorito()

	return boot.Xorriso.Build(w.GetDirs().Iso, w.target.Out)
})
