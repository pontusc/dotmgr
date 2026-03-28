package status

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd returns the status subcommand.
func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the state of managed dotfiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("status: not yet implemented")
			return nil
		},
	}
}