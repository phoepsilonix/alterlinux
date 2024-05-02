package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/check"
	"github.com/Hayao0819/nahi/cobrautils"
)

func init() {
	cobrautils.AddSubCmds(check.Cmd())
}
