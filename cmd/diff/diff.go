package diff

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd returns the diff subcommand.
func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Show differences between local files and their configured sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("diff: not yet implemented")
			return nil
		},
	}
}