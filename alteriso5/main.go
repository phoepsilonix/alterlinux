package main

import (
	"os"

	"github.com/FascodeNet/alterlinux/alteriso5/cmd"
)

func main() {
	root := cmd.Root()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
