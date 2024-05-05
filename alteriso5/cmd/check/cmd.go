package check

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func check() error {
	if os.Getuid() != 0 {
		return errors.New("this program must be run as root")
	}
	return nil
}

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "check",
		Short: "Check if the system is ready to build the ISO",
		SilenceUsage: true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return check()
		},
	}

	return &cmd
}
