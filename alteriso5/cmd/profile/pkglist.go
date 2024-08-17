package profile

import (
	"slices"

	"github.com/FascodeNet/alterlinux/alteriso5/config/pkg"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

func pkgListCmd() *cobra.Command {
	showContent := false
	cmd := cobra.Command{
		Use:   "pkglist profile arch",
		Short: "List packages in the profile",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			files, err := pkg.FindPkgListFiles(profile.Base, args[1])
			if err != nil {
				return err
			}

			if showContent {
				pkglist := []string{}
				for _, f := range files {
					pkgs, err := pkg.ReadPkgListFile(f)
					if err != nil {
						return err
					}
					pkglist = append(pkglist, pkgs...)
				}
				slices.Sort(pkglist)
				funk.ForEach(pkglist, func(p string) {
					cmd.Println(p)
				})

			} else {
				funk.ForEach(files, func(f string) {
					cmd.Println(f)
				})
			}

			return nil

		},
	}

	cmd.Flags().BoolVarP(&showContent, "content", "c", false, "Show content of the package list file")

	return &cmd
}

func init() {
	subCmds.Add(pkgListCmd())
}
