package cmd

import (
	"fmt"
	"os"

	"github.com/smanierre/opncli/cmd/core"
	"github.com/smanierre/opncli/internal/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "opnsense-cli",
		Short: "CLI to interact with OPNSense api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.Cfg = config.Init(cfgFile)

	})

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error determining users home directory: %s\n", err.Error())
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", fmt.Sprintf("%s/.config/opnsense-cli/config.json", home), "Config file (JSON) to be used for cli.")

	rootCmd.AddCommand(core.CoreCommand)
}
