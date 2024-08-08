package unbound

import (
	"github.com/smanierre/opncli/cmd/core/unbound/host_alias"
	"github.com/smanierre/opncli/cmd/core/unbound/host_overrides"

	"github.com/spf13/cobra"
)

func init() {
	UnboundCommand.AddCommand(host_overrides.HostOverrides)
	UnboundCommand.AddCommand(host_alias.HostAliases)
}

var UnboundCommand = &cobra.Command{
	Use:   "unbound",
	Short: "Manage Unbound DNS settings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
