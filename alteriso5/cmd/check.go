package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/check"
)

func init() {
	rootSubCmds.Add(check.Cmd())
}
