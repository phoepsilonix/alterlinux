package build

import (
	"fmt"
	"os"
	"path"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/config"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/build/work"
	"github.com/FascodeNet/alterlinux/alteriso5/cmd/check"
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "build",
		Short: "Build an ISO image",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Handle signals
			utils.OnSignal(func(s os.Signal) {
				fmt.Println("Received signal:", s)
				os.Exit(1)
			}, os.Interrupt)

			// Check dependencies
			if err := utils.CallCmd(cmd, *check.Cmd()); err != nil {
				return err
			}

			// Get current working directory
			current, err := os.Getwd()
			if err != nil {
				return err
			}
			workDir := path.Join(current, "work")
			outDir := path.Join(current, "out")

			// Prepare work environment
			work, err := work.New(workDir)
			if err != nil {
				return err
			}

			// Dummy profile
			profile := config.Profile{
				Base:       path.Join(current, "profile"),
				InstallDir: "alter",
				BootModes:  []string{"SysLinux"},
			}

			// TODO: Add more targets
			target := config.NewTarget("x86_64", outDir)
			return work.Build(profile, target, cmd)

		},
	}

	return &cmd
}
