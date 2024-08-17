package profile

import (
	"errors"

	"github.com/spf13/cobra"
)

func convertCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "convert profile out",
		Short: "Convert profile for archiso",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return &cmd

}

func init() {
	subCmds.Add(convertCmd())
}
