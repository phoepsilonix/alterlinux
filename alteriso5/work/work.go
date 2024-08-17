package work

import (
	"log/slog"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/work/airootfs"
	osutils "github.com/Hayao0819/nahi/futils"
	"github.com/spf13/cobra"
)

type Work struct {
	Base    string
	profile *config.Profile
	target  *config.Target
	Cmd     *cobra.Command
	Dirs    *dirs
	Files   *files
}

// Workで使うディレクトリのパスをまとめた構造体
type dirs struct {
	Current  string
	Data     string
	Work     string
	Pacstrap string
	Iso      string
	WorkArch string
	Efiboot  string
}

// Workで使うファイルのパスをまとめた構造体
type files struct {
	EfibootImg string
}

type configValues struct {
	Arch  string
	Label string
}

func New(dir string) *Work {
	w := Work{
		Base: dir,
	}
	return &w
}

func (w *Work) GetDirs() (*dirs, error) {
	current, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dv := dirs{
		Current:  current,
		Data:     path.Join(current, "alteriso5", "data"),
		Work:     w.Base,
		Pacstrap: path.Join(w.Base, w.target.Arch, "airootfs"),
		Iso:      path.Join(w.Base, "iso"),
		WorkArch: path.Join(w.Base, w.target.Arch),
		Efiboot:  path.Join(w.Base, w.target.Arch, "efiboot"),
	}

	return &dv, nil
}

func (w *Work) GetFiles() (*files, error) {
	fv := files{
		EfibootImg: path.Join(w.Dirs.Work, "efiboot.img"),
	}

	return &fv, nil
}

func (w *Work) GetChroot() (*airootfs.Chroot, error) {
	return airootfs.GetChrootDir(w.Dirs.Pacstrap, w.target.Arch, path.Join(w.profile.Base, "pacman.conf"))
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

func (w *Work) Values() *configValues {
	v := configValues{
		Arch:  w.target.Arch,
		Label: w.profile.ISOLabel,
	}

	return &v
}
