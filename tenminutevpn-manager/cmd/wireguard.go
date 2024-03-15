package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func wireguardSetup() {
	iface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := wireguard.NewWireguard("wg0", iface, "100.96.0.1/24", 51820)
	if err != nil {
		log.Fatal(err)
		return
	}

	peer1, err := wireguard.NewWireguard("peer-1", "", "100.96.0.2/32", 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	server.AddPeer(peer1)

	err = server.WriteConfig("/tmp/wg0.conf")
	if err != nil {
		log.Fatal(err)
		return
	}

}

var wireguardSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup WireGuard Interface",
	Run: func(cmd *cobra.Command, args []string) {
		wireguardSetup()
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
