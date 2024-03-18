package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

func wireguardSetup(configPath, peerPath string) {
	iface, err := network.GetDefaultInterface()
	if err != nil {
		log.Fatal(err)
		return
	}

	wg, err := wireguard.NewWireguard("wg0", iface, "100.96.0.1/24", 51820)
	if err != nil {
		log.Fatal(err)
		return
	}

	peer, err := wireguard.NewWireguard("peer1", "", "100.96.0.2/32", 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	wg.AddPeer(peer)

	err = wg.Write(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = peer.Write(peerPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wg.SystemdService().Enable()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wg.SystemdService().Start()
	if err != nil {
		log.Fatal(err)
		return
	}
}

var wireguardSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup WireGuard Interface",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
		peerPath := cmd.Flag("peer").Value.String()
		wireguardSetup(configPath, peerPath)
	},
}

var wireguardCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "TenMinuteVPN Wireguard",
}

func init() {
	wireguardSetupCmd.Flags().StringP("config", "c", "/etc/wireguard/wg0.conf", "Path to WireGuard configuration file")
	wireguardSetupCmd.Flags().StringP("peer", "p", "/etc/wireguard/peer-1.conf", "Path to WireGuard peer configuration file")

	wireguardCmd.AddCommand(wireguardSetupCmd)
	rootCmd.AddCommand(wireguardCmd)
}
