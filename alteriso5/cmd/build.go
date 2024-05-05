package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build"
)

func init() {
	rootSubCmds.RegisterSubCmd(build.Cmd())
}
