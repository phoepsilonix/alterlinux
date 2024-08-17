package profile

import (
	"github.com/FascodeNet/alterlinux/alteriso5/config/pkg"
	"github.com/spf13/cobra"
)

func pkgListCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pkglist profile arch",
		Short: "List packages in the profile",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			files, err := pkg.FindPkgListFiles(profile.Base, args[1])
			if err != nil{
				return err
			}

			cmd.Println(files)
			return nil
			
		},
	}

	return &cmd
}

func init() {
	subCmds.Add(pkgListCmd())
}
