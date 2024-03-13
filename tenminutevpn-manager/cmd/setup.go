package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
)

func setup() {
	iface, err := network.GetDefaultInterface()
	if err != nil {
		fmt.Println("failed to get default interface:", err)
		return
	}
	fmt.Println(iface)

	privateip, err := network.GetPrivateIPv4(iface)
	if err != nil {
		fmt.Println("failed to get private ip:", err)
		return
	}
	fmt.Println(privateip.String())

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
