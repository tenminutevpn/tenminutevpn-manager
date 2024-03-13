package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard"
)

var wireguardGeneratePrivateKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate a Wireguard Keypair",
	Run: func(cmd *cobra.Command, args []string) {
		privkey, err := wireguard.GenKey()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(privkey)
	},
}

var wireguardGeneratePublicKeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Generate a Wireguard Public Key",
	Run: func(cmd *cobra.Command, args []string) {
		pubkey, err := wireguard.GenPublicKey(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(pubkey)
	},
}

func wireguardSetup(cmd *cobra.Command, args []string) {
	fmt.Println("setup called")
}

var wireguardSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup WireGuard Interface",
	Run:   wireguardSetup,
}

var wireguardCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "TenMinuteVPN Wireguard",
}

func init() {
	rootCmd.AddCommand(wireguardCmd)

	wireguardCmd.AddCommand(wireguardGeneratePrivateKeyCmd)
	wireguardCmd.AddCommand(wireguardGeneratePublicKeyCmd)
	wireguardCmd.AddCommand(wireguardSetupCmd)
}
