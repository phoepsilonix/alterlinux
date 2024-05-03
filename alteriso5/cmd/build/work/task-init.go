package work

import (
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
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
