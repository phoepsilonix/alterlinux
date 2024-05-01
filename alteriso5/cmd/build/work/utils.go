package work

import (
	"path"

	"github.com/Hayao0819/nahi/osutils"
)

func (w *Work) RunOnce(task *BuildTask) error {
	lp := path.Join(w.Base, w.target.Arch, "lockfile", task.Name())
	if osutils.Exists(lp) {
		return nil
	}

	return task.Run(w)
}
