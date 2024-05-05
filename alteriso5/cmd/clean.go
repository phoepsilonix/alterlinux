package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/clean"
)

func init() {
	rootSubCmds.RegisterSubCmd(clean.Cmd())
}
