package work

import (
	"log/slog"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/work/airootfs"
	"github.com/Hayao0819/nahi/osutils"
	"github.com/spf13/cobra"
)

type Work struct {
	Base    string
	profile *config.Profile
	target  *config.Target
	Cmd     *cobra.Command
}

type dirs struct {
	Work           string
	Pacstrap       string
	Iso            string
	SyslinuxConfig string
}

func New(dir string) (*Work, error) {
	return &Work{
		Base: dir,
	}, nil
}

func (w *Work) GetDirs() *dirs {
	return &dirs{
		Work:           w.Base,
		Pacstrap:       path.Join(w.Base, w.target.Arch, "airootfs"),
		Iso:            path.Join(w.Base, "iso"),
		SyslinuxConfig: path.Join(w.profile.Base, "syslinux"),
	}
}

func (w *Work) GetChroot() (*airootfs.Chroot, error) {
	return airootfs.GetChrootDir(w.GetDirs().Pacstrap, w.target.Arch, path.Join(w.profile.Base, "pacman.conf"))
}

func (w *Work) RunOnce(task *BuildTask) error {
	lp := path.Join(w.Base, w.target.Arch, "lockfile", task.Name())
	if osutils.Exists(lp) {
		slog.Warn("This task has already runned", "name", task.Name())
		return nil
	}

	if err := task.Run(w); err != nil {
		return err
	} else {
		// Dont care about error
		os.MkdirAll(path.Dir(lp), 0755)
		os.Create(lp)

	}
	return nil
}
