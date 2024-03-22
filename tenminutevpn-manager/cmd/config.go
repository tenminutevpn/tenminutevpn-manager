package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		if err := config.ParseResources(configFile); err != nil {
			panic(err)
		}
	},
}

func init() {
	configCmd.Flags().StringP("file", "f", "/etc/tenminutevpn/tenminutevpn.yaml", "config file")

	rootCmd.AddCommand(configCmd)
}
