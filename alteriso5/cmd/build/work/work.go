package work

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/chroot"
	"github.com/spf13/cobra"
)

type BuildTask struct {
	name string
	task func(work *Work) error
}

func NewBuildTask(name string, task func(*Work) error) *BuildTask {
	return &BuildTask{
		name: name,
		task: task,
	}
}

func (t *BuildTask) Name() string {
	return t.name
}

func (t *BuildTask) Run(w *Work) error {
	return t.task(w)
}

type Work struct {
	Base    string
	Chroots []*chroot.Env
	profile *config.Profile
	target  *config.Target
	Cmd     *cobra.Command
}
