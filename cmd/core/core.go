package core

import (
	"github.com/smanierre/opncli/cmd/core/unbound"

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
