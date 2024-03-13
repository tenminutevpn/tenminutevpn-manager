package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
)

func setup() {
	ip, err := network.GetPublicIPv4()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip.String())
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
