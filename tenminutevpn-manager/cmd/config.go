package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := config.ParseConfig()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
