package work

import (
	"log/slog"
	"path"

	"github.com/Hayao0819/nahi/osutils"
	cp "github.com/otiai10/copy"
)

var makeBaseDirs *BuildTask = NewBuildTask("makeBaseDirs",
	func(work Work) error {

		dirs := []string{
			work.Base,
			work.target.Out,
		}

		if err := osutils.MkdirsAll(dirs...); err != nil {
			return err
		}
		return nil
	})

var makeCustomAirootfs *BuildTask = NewBuildTask("makeCustomAirootfs", func(w Work) error {
	slog.Info("Copying custom airootfs files...")
	profileAirootfsDir := path.Join(w.profile.Base, "airootfs")

	if err := cp.Copy(profileAirootfsDir, w.Dirs.Pacstrap, cp.Options{
		PreserveOwner: true,
		Sync:          true,
	}); err != nil {
		return err
	}

	return nil
})

var makeChroot *BuildTask = NewBuildTask("makeChroot", func(work Work) error {

	env, err := work.GetChroot()
	if err != nil {
		return err
	}

	pkglist, err := work.profile.GetPkgList(work.target.Arch)
	if err != nil {
		return err
	}

	if err := env.Init(pkglist...); err != nil {
		return err

	}
	return nil
})
