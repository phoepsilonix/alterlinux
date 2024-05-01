package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/clean"
	"github.com/Hayao0819/nahi/cobrautils"
)

func init() {
	cobrautils.AddSubCmds(clean.Cmd())
}
