package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/airootfs"
)

var makeAirootfs *BuildTask = NewBuildTask("makeAirootfs", func(w *Work) error {

	slog.Debug("Copying profile to airootfs...")
	airootfsDir := path.Join(w.Base, w.target.Arch, "airootfs")

	sqfs := airootfs.SquashFS{
		Base: airootfsDir,
		Out:  path.Join(w.GetDirs().Iso, w.profile.InstallDir, w.target.Arch, "airootfs.sfs"),
	}

	return sqfs.Build()
})
