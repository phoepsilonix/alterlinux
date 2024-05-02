package utils

import "github.com/spf13/cobra"

func CallCmd(me *cobra.Command, target cobra.Command, args ...string) error {
	target.SetOut(me.OutOrStdout())
	target.SetErr(me.OutOrStderr())
	target.SetIn(me.InOrStdin())
	target.SetArgs(args)
	return target.Execute()
}
