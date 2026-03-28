package sync

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd returns the sync subcommand.
func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Sync dotfiles from configured sources to local paths",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("sync: not yet implemented")
			return nil
		},
	}
}