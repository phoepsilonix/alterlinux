package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := cobra.Command{
		Use:   "alteriso5",
		Short: "AlterISO5 is a tool to build Arch Linux live ISO images",
	}

	cobrautils.AddSubCmdsToRoot(&root)

	return &root
}
