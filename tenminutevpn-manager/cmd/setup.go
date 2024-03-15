package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func setup() {
	wgAddress := "100.96.0.1/24"
	wgName := "wg0"
	wgPort := 51820

	ip, ipNet, err := net.ParseCIDR(wgAddress)
	if err != nil {
		log.Fatalf("failed to parse CIDR: %s", err.Error())
		return
	}

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

	serverConfig := wg.ToServerConfig()
	err = serverConfig.WriteToFile(fmt.Sprintf("/tmp/%s.conf", wgName))
	if err != nil {
		log.Fatalf("failed to write server config: %s", err.Error())
		return
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
