package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func setup() {
	fmt.Println("Setting up TenMinuteVPN")
	wireguardSetup("/etc/wireguard/wg0.conf", "/etc/wireguard/peer-1.conf")
	squidSetup(3128)
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
