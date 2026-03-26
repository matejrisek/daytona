package profile

import (
	"github.com/daytonaio/daytona/cli/internal"
	"github.com/spf13/cobra"
)

func InitProfileCommands(root *cobra.Command) {
	rootProfileCmd := &cobra.Command{
		Use:     "profile",
		Short:   "Manage Daytona profiles",
		Args:    cobra.NoArgs,
		GroupID: internal.USER_GROUP,
		RunE:    rootCommand,
	}

	root.AddCommand(rootProfileCmd)
}
