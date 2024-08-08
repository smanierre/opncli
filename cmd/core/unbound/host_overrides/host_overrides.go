package host_overrides

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/smanierre/opncli/api"
	"github.com/smanierre/opncli/internal/client"
	"github.com/smanierre/opncli/internal/config"

	"github.com/spf13/cobra"
)

func init() {
	HostOverrides.AddCommand(getOverrides)
	HostOverrides.AddCommand(listOverrides)
	HostOverrides.AddCommand(addOverride)
	HostOverrides.AddCommand(removeCommand)

	// Add command flags
	addOverride.Flags().StringVar(&addEnabled, "enabled", "1", "Whether or not the override is enabled")
	addOverride.Flags().StringVar(&addDescription, "description", "", "Description for the host override")
	addOverride.Flags().StringVar(&addDomain, "domain", "", "Domain for the host override")
	addOverride.Flags().StringVar(&addHostname, "hostname", "", "Hostname for the override")
	addOverride.Flags().StringVar(&addMx, "mx", "", "Whether the record should be an MX record or not")
	addOverride.Flags().StringVar(&addMxprio, "mxprio", "", "Priority if an MX record")
	addOverride.Flags().StringVar(&addRecordType, "record_type", "A", "Type of record")
	addOverride.Flags().StringVar(&addIpAddress, "ip", "", "Ip address of host")
	addOverride.MarkFlagRequired("domain")
	addOverride.MarkFlagsRequiredTogether("mx", "mxprio")
	addOverride.MarkFlagsRequiredTogether("ip", "hostname")
}

var HostOverrides = &cobra.Command{
	Use:   "host-overrides",
	Short: "Manage Unbound DNS host overrides",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	addEnabled     string
	addDescription string
	addDomain      string
	addHostname    string
	addMx          string
	addMxprio      string
	addRecordType  string
	addIpAddress   string
)

var addOverride = &cobra.Command{
	Use:   "add",
	Short: "Add an Unbound DNS host override",
	RunE: func(cmd *cobra.Command, args []string) error {
		override := api.AddHostOverride{
			Description: addDescription,
			Domain:      addDomain,
			Enabled:     addEnabled,
			HostName:    addHostname,
			Mx:          addMx,
			MxPriority:  addMxprio,
			RecordType:  addRecordType,
			IpAddress:   addIpAddress,
		}
		fmt.Println(override.String())
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodPost, "unbound", "settings", "addHostOverride", strings.NewReader(override.String()))
		if err != nil {
			return err
		}
		var r api.AddItemRes
		err = json.NewDecoder(body).Decode(&r)
		if err != nil {
			return err
		}
		fmt.Print(r)
		return nil
	},
}

var getOverrides = &cobra.Command{
	Use:   "get",
	Short: "Get Unbound DNS host overrides",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SetUsageFunc(func(c *cobra.Command) error {
			fmt.Println("Usage:")
			fmt.Println(c.CommandPath(), "<override uuid>")
			return nil
		})

		valid := cobra.ExactArgs(1)
		err := valid(cmd, args)
		if err != nil {
			return err
		}

		c := client.New(config.Cfg)
		body, err := c.PerformRequest(http.MethodGet, "unbound", "settings", "getHostOverride", nil, args[0])
		if err != nil {
			return err
		}
		var override api.SingleHostOverrideRes
		err = json.NewDecoder(body).Decode(&override)
		if err != nil {
			return err
		}
		fmt.Print(override.HostOverride)
		return nil
	},
}

var listOverrides = &cobra.Command{
	Use:   "list",
	Short: "List all DNS overrides",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New(config.Cfg)
		body, err := c.PerformRequest(http.MethodGet, "unbound", "settings", "searchHostOverride", nil)
		if err != nil {
			return err
		}
		var overrides api.HostOverrideList
		err = json.NewDecoder(body).Decode(&overrides)
		if err != nil {
			return err
		}
		fmt.Print(overrides)
		return err
	},
}

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a host override",
	RunE: func(cmd *cobra.Command, args []string) error {
		checkArgs := cobra.ExactArgs(1)
		err := checkArgs(cmd, args)
		if err != nil {
			return err
		}
		client := client.New(config.Cfg)
		body, err := client.PerformRequest(http.MethodPost, "unbound", "settings", "delHostOverride", nil, args[0])
		if err != nil {
			return err
		}
		var res api.DeleteItemRes
		if err = json.NewDecoder(body).Decode(&res); err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}
