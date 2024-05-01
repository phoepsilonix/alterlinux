package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build"
	"github.com/Hayao0819/nahi/cobrautils"
)

func init() {
	cobrautils.AddSubCmds(build.Cmd())
}
