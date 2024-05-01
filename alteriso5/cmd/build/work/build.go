package work

import "github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"

func (work Work) Build(p config.Profile, t config.Target) error {

	work.profile = &p
	work.target = &t

	tasks := []*BuildTask{
		makeBaseDirs,
		makeChroot,
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
