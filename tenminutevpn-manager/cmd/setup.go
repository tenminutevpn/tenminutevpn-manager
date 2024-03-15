package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func setup() {
	ip, ipNet, err := net.ParseCIDR("100.96.0.1/24")
	if err != nil {
		log.Fatalf("failed to parse CIDR: %s", err.Error())
		return
	}

	wgName := "wg0"

	iface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	wg := wireguard.NewWireguard(iface, wgName, ip, ipNet.Mask, 51820)

	wgPrivateKey, err := wireguard.GenKey()
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err.Error())
		return
	}
	wg.SetPrivateKey(wgPrivateKey)

	log.Println(wg.RenderConfig())

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
