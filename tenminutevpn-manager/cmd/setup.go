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
	wgPort := 51820

	iface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	wg := wireguard.NewWireguard(wgName, iface, ip, ipNet, wgPort)

	wgPrivateKey, err := wireguard.GenKey()
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err.Error())
		return
	}
	wg.SetPrivateKey(wgPrivateKey)

	data := wg.GetTemplateData()
	log.Println(wg.RenderConfig(data))

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
