package work

import (
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/chroot"
)

func (work *Work) MakeBaseDirs() error {

	dirs := []string{
		work.Base,
		work.target.Out,
	}

	if err := utils.MkdirsAll(dirs...); err != nil {
		return err
	}
	return nil
}

func (work *Work) MakeChroot() error {
	for _, arch := range work.target.Arch {
		dir := path.Join(work.Base, arch)
		env := chroot.New(dir, arch)
		if err := env.Init(); err != nil {
			return err
		}
	}
	return nil
}
