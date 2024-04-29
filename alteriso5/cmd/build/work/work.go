package work

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work/chroot"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
)

type BuildTask func() error

type Work struct {
	Base    string
	Chroots []*chroot.Env
	profile *config.Profile
	target  *config.Target
}
