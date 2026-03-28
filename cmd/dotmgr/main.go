package main

import (
	"fmt"
	"os"

	"dotmgr/cmd/diff"
	"dotmgr/cmd/status"
	"dotmgr/cmd/sync"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "dotmgr",
		Short: "Stateless dotfiles manager",
		Long:  "dotmgr retrieves dotfiles from git repositories and places them at configured paths.",
	}

	root.AddCommand(sync.Cmd())
	root.AddCommand(diff.Cmd())
	root.AddCommand(status.Cmd())

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}