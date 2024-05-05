package clean

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "clean",
		Short: "Clean the working directory",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return &cmd
	
}
