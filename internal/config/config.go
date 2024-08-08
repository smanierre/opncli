package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cfg = Config{}

type Config struct {
	Host      string `json:"host"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func Init(cfgFile string) Config {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgFile = fmt.Sprintf("%s/.config/opnsense-cli/config.json", home)
	}

	var config = Config{}
	f, err := os.Open(cfgFile)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		err = generateConfig(&config)
		cobra.CheckErr(err)
		parts := strings.Split(cfgFile, "/")
		configDir := strings.Join(parts[:len(parts)-1], "/")
		err = os.MkdirAll(configDir, 0750)
		cobra.CheckErr(err)
		f, err = os.Create(cfgFile)
		cobra.CheckErr(err)
		err = json.NewEncoder(f).Encode(config)
		cobra.CheckErr(err)
	} else {
		cobra.CheckErr(err)
		err = json.NewDecoder(f).Decode(&config)
		cobra.CheckErr(err)
	}
	return config
}

func generateConfig(config *Config) error {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("No config file found, generating new config")
	fmt.Print("OPNSense hostname: ")
	chars, _, err := input.ReadLine()
	if err != nil {
		return err
	}
	config.Host = string(chars)
	fmt.Print("API Key: ")
	chars, _, err = input.ReadLine()
	if err != nil {
		return err
	}
	config.ApiKey = string(chars)
	fmt.Print("API Secret: ")
	chars, _, err = input.ReadLine()
	if err != nil {
		return err
	}
	config.ApiSecret = string(chars)
	return nil
}
