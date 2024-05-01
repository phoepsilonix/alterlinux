package build

import (
	"fmt"
	"os"
	"time"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "build",
		Short: "Build an ISO image",
		RunE: func(cmd *cobra.Command, args []string) error {

			utils.OnSignal(func(s os.Signal) {
				fmt.Println("Received signal:", s)
				os.Exit(1)
			}, os.Interrupt)
			time.Sleep(20 * time.Second)

			if err := check(); err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			if err := build(); err != nil {
				cmd.PrintErrln(err)
			}

			return nil
		},
	}

	return &cmd
}
