package work

import (
	"path"

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
