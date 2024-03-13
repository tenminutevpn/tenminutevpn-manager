package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func wireguardGenerateKeypair(cmd *cobra.Command, args []string) {
	fmt.Println("generate-keypair called")
}

var wireguardGenerateKeypairCmd = &cobra.Command{
	Use:   "generate-keypair",
	Short: "Generate a Wireguard Keypair",
	Long:  "Generate a Wireguard Keypair",
	Run:   wireguardGenerateKeypair,
}

var wireguardCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "TenMinuteVPN Wireguard",
	Long:  "Wireguard is a fast, modern, and secure VPN tunnel",
}

func init() {
	rootCmd.AddCommand(wireguardCmd)

	wireguardCmd.AddCommand(wireguardGenerateKeypairCmd)
}
