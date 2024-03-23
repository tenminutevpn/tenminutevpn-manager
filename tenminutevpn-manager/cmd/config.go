package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
		}

		if err := config.Parse(configPath); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	configCmd.Flags().StringP("file", "f", "/etc/tenminutevpn/tenminutevpn.yaml", "config file")

	rootCmd.AddCommand(configCmd)
}
