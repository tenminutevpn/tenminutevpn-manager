package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// generate a new Wireguard private key using the `wg` command
func wireguardGeneratePrivateKey(_ *cobra.Command, args []string) {
	cmd := exec.Command("wg", "genkey")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

var wireguardGeneratePrivateKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate a Wireguard Keypair",
	Run:   wireguardGeneratePrivateKey,
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
	wireguardCmd.AddCommand(wireguardSetupCmd)
}
