package cmd

import (
	"github.com/FascodeNet/alterlinux/alteriso5/log"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var rootSubCmds = cobrautils.Registory{}

func Root() *cobra.Command {
	root := cobra.Command{
		Use:   "alteriso5",
		Short: "AlterISO5 is a tool to build Arch Linux live ISO images",
		PersistentPreRunE: cobrautils.WithParentPersistentPreRunE(func(cmd *cobra.Command, args []string) error {
			log.Setup()
			return nil
		}),
		SilenceUsage: true,
	}

	rootSubCmds.BindSubCmds(&root)

	return &root
}
