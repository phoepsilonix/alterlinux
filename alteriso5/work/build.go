package work

import (
	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/spf13/cobra"
)

func (work Work) Build(p config.Profile, t config.Target, c *cobra.Command) error {

	// Setup work
	work.profile = &p
	work.target = &t
	work.Cmd = c
	dirs, err := work.GetDirs()
	if err != nil {
		return err
	}
	work.Dirs = dirs

	files, err := work.GetFiles()
	if err != nil {
		return err
	}
	work.Files = files

	tasks := []*BuildTask{
		validate,
		makeBaseDirs,
		makeCustomAirootfs,
		makeChroot,
		makeBootModes,
		makeAirootfs,
		makeOutFiles,
	}

	for _, t := range tasks {
		if err := work.RunOnce(t); err != nil {
			return err
		}
	}

	return nil
}
