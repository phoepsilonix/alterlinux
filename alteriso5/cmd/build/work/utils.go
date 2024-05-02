package work

import (
	"os"
	"path"

	"github.com/Hayao0819/nahi/osutils"
)

func (w *Work) RunOnce(task *BuildTask) error {
	lp := path.Join(w.Base, w.target.Arch, "lockfile", task.Name())
	if osutils.Exists(lp) {
		return nil
	}

	if err := task.Run(w); err != nil {
		return err
	} else {
		// Dont care about error
		os.Create(lp)

	}
	return nil
}
