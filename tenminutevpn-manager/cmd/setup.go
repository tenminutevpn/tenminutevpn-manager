package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
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

	privkey, pubkey, err := wireguard.GenKeypair()
	if err != nil {
		log.Fatalf("failed to generate keypair: %w", err)
	}

	folder := "/etc/wireguard"
	err = wireguard.WriteKeypair(folder, privkey, pubkey)
	if err != nil {
		log.Fatalf("failed to write keypair: %w", err)
	}

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
