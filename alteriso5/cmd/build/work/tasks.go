package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/airootfs"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/boot"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/chroot"
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	cp "github.com/otiai10/copy"
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

		isoSyslinuxDir := path.Join(w.Base, "iso", "boot", "syslinux")

		if err := utils.MkdirsAll(isoSyslinuxDir); err != nil {
			return err
		}

		profileSysLinuxDir := path.Join(w.profile.Base, "syslinux")

		if err := cp.Copy(profileSysLinuxDir, isoSyslinuxDir); err != nil {
			return err
		}

		boot.Xorriso.SetArgsForSysLinuxElTorito()
		boot.Xorriso.SetArgsForSysLinuxElTorito()

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
	return boot.Xorriso.Build(w.Base, w.target.Out)
})
