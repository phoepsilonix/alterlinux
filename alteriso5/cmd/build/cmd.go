package build

import (
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/check"
	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/FascodeNet/alterlinux/alteriso5/work"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "build",
		Short: "Build an ISO image",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Handle signals
			utils.OnSignal(func(s os.Signal) {
				cmd.Println("Received signal:", s)
				os.Exit(1)
			}, os.Interrupt)

			// Check dependencies
			if err := cobrautils.CallCmd(cmd, *check.Cmd()); err != nil {
				return err
			}

			// Get current working directory
			current, err := os.Getwd()
			if err != nil {
				return err
			}
			workDir := path.Join(current, "work")
			outDir := path.Join(current, "out")

			// Read profile
			profile, err := config.ReadProfile(args[0])
			if err != nil {
				return err
			}

			// TODO: Add more targets
			target := config.NewTarget("x86_64", outDir)
			return work.New(workDir).Build(*profile, target, cmd)

		},
	}

	return &cmd
}
