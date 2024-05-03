package work

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
	"github.com/spf13/cobra"
)

func (work Work) Build(p config.Profile, t config.Target, c *cobra.Command) error {

	work.profile = &p
	work.target = &t
	work.Cmd = c

	tasks := []*BuildTask{
		makeBaseDirs,
		makeChroot,
		makeBootModes,
		makeAirootfs,
		makeOutFiles,
	}

	for _, t := range tasks {
		//err := (*t).task(&work)

		err := work.RunOnce(t)
		if err != nil {
			return err
		}
	}

	return nil
}
