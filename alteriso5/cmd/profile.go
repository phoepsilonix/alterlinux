package cmd

import "github.com/FascodeNet/alterlinux/alteriso5/cmd/profile"

func init() {
	rootSubCmds.RegisterSubCmd(profile.Cmd())
}
