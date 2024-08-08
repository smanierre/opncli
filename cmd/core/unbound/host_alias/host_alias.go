package host_alias

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opnsense-cli/internal/api"
	"opnsense-cli/internal/client"
	"opnsense-cli/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	HostAliases.AddCommand(listAll)
	HostAliases.AddCommand(getAlias)
	HostAliases.AddCommand(addAlias)
	HostAliases.AddCommand(removeAlias)

	//Add command flags
	addAlias.Flags().StringVar(&addEnabled, "enabled", "1", "Whether or not the alias should be active")
	addAlias.Flags().StringVar(&addDomain, "domain", "", "Domain to use for the alias")
	addAlias.Flags().StringVar(&addDescription, "description", "", "Description for the alias")
	addAlias.Flags().StringVar(&addHostOverride, "host-override", "", "UUID of the host override the alias should be assigned to")
	addAlias.Flags().StringVar(&addHostName, "alias", "", "The alias you want to point to the host override")
	addAlias.MarkFlagRequired("domain")
	addAlias.MarkFlagRequired("host-override")
	addAlias.MarkFlagRequired("alias")
}

var HostAliases = &cobra.Command{
	Use:   "host-aliases",
	Short: "Manage Unbound DNS host aliases",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listAll = &cobra.Command{
	Use:   "list",
	Short: "List all host aliases",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodGet, "unbound", "settings", "searchHostAlias", nil)
		if err != nil {
			return err
		}
		var aliases api.HostAliasList
		err = json.NewDecoder(body).Decode(&aliases)
		if err != nil {
			return err
		}
		fmt.Print(aliases)
		return nil
	},
}

var getAlias = &cobra.Command{
	Use:   "get",
	Short: "Get a single host alias",
	RunE: func(cmd *cobra.Command, args []string) error {
		checkArgs := cobra.ExactArgs(1)
		err := checkArgs(cmd, args)
		if err != nil {
			return err
		}
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodGet, "unbound", "settings", "getHostAlias", nil, args[0])
		if err != nil {
			return err
		}
		var res api.SingleHostAliasRes
		err = json.NewDecoder(body).Decode(&res)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

var (
	addEnabled      string
	addDomain       string
	addDescription  string
	addHostOverride string
	addHostName     string
)

var addAlias = &cobra.Command{
	Use:   "add",
	Short: "Add a host alias for a given host override",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := api.AddHostAlias{
			Description:    addDescription,
			Domain:         addDomain,
			Enabled:        addEnabled,
			HostOverrideID: addHostOverride,
			Alias:          addHostName,
		}
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodPost, "unbound", "settings", "addHostAlias", strings.NewReader(a.String()))
		if err != nil {
			return err
		}
		var res api.AddItemRes
		err = json.NewDecoder(body).Decode(&res)
		if err != nil {
			return err
		}
		err = client.ReconfigureUnbound()
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

var removeAlias = &cobra.Command{
	Use:   "remove",
	Short: "Remove a host alias",
	RunE: func(cmd *cobra.Command, args []string) error {
		checkArgs := cobra.ExactArgs(1)
		err := checkArgs(cmd, args)
		if err != nil {
			return err
		}
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodPost, "unbound", "settings", "delHostAlias", nil, args[0])
		if err != nil {
			return err
		}
		var res api.DeleteItemRes
		err = json.NewDecoder(body).Decode(&res)
		if err != nil {
			return err
		}
		if err = client.ReconfigureUnbound(); err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}
