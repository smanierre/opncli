package core

import (
	"opnsense-cli/cmd/core/unbound"

	"github.com/spf13/cobra"
)

func init() {
	CoreCommand.AddCommand(unbound.UnboundCommand)
}

var CoreCommand = &cobra.Command{
	Use:   "core",
	Short: "Core OPNSense services",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
