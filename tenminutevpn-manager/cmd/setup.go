package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func setup() {
	netIface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	wg := wireguard.NewWireguard(netIface, "wg0", nil, 51820)

	wgPrivateKey, err := wireguard.GenKey()
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err.Error())
		return
	}

	wg.SetPrivateKey(wgPrivateKey)

	privateKey, _, err := wireguard.GenKeypair()
	if err != nil {
		log.Fatalf("failed to generate keypair: %s", err.Error())
	}

	// folder := "/etc/wireguard"
	// err = wireguard.WriteKeypair(folder, privateKey, pubkey)
	// if err != nil {
	// 	log.Fatalf("failed to write keypair: %s", err.Error())
	// }

	iface, err := network.GetDefaultInterface()
	if err != nil {
		fmt.Println("failed to get default interface:", err)
		return
	}
	fmt.Println(iface)

	config := wireguard.GenServerConfig(iface, privateKey)
	fmt.Println(config)

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
