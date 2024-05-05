package utils

import "github.com/spf13/cobra"

func CallCmd(me *cobra.Command, target cobra.Command, args ...string) error {
	target.SetOut(me.OutOrStdout())
	target.SetErr(me.OutOrStderr())
	target.SetIn(me.InOrStdin())
	target.SetArgs(args)
	return target.Execute()
}

func WithParentPersistentPreRunE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		parent := cmd.Parent()
		if parent != nil {
			if parent.PersistentPreRunE != nil {
				err := parent.PersistentPreRunE(parent, args)
				if err != nil {
					return err
				}
			} else if parent.PersistentPreRun != nil {
				parent.PersistentPreRun(parent, args)
			}
		}

		return f(cmd, args)
	}
}
