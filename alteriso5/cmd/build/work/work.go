package work

import (
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
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

func (w *Work) GetDirs() *dirs {
	return &dirs{
		Work:           w.Base,
		Pacstrap:       path.Join(w.Base, w.target.Arch, "airootfs"),
		Iso:            path.Join(w.Base, "iso"),
		SyslinuxConfig: path.Join(w.profile.Base, "syslinux"),
	}
}
