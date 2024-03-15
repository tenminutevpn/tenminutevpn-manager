package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func wireguardSetup(wgName string, wgAddress string, wgPort int) {
	iface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	wg, err := wireguard.NewWireguard(wgName, iface, wgAddress, wgPort)
	if err != nil {
		log.Fatalf("failed to create wireguard: %s", err.Error())
		return
	}

	err = wg.WriteServerConfig(fmt.Sprintf("/tmp/%s.conf", wgName))
	if err != nil {
		log.Fatalf("failed to write server config: %s", err.Error())
		return
	}
}

var wireguardSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup WireGuard Interface",
	Run: func(cmd *cobra.Command, args []string) {
		wireguardSetup("wg0", "100.96.0.1/24", 51820)
	},
}

var wireguardCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "TenMinuteVPN Wireguard",
}

func init() {
	rootCmd.AddCommand(wireguardCmd)

	wireguardCmd.AddCommand(wireguardSetupCmd)
}
