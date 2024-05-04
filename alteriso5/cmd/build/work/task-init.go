package work

import (
	"log/slog"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	cp "github.com/otiai10/copy"
)

var makeBaseDirs *BuildTask = NewBuildTask("makeBaseDirs",
	func(work Work) error {

		dirs := []string{
			work.Base,
			work.target.Out,
		}

		if err := utils.MkdirsAll(dirs...); err != nil {
			return err
		}
		return nil
	})

var makeCustomAirootfs *BuildTask = NewBuildTask("makeCustomAirootfs", func(w Work) error {
	slog.Info("Copying custom airootfs files...")
	profileAirootfsDir := path.Join(w.profile.Base, "airootfs")

	if err := cp.Copy(profileAirootfsDir, w.GetDirs().Pacstrap, cp.Options{
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
	if err := env.Init(); err != nil {
		return err

	}
	return nil
})
