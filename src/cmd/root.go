package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
	Use:   "tenminutevpn-manager",
	Short: "TenMinuteVPN Manager",
	Long: "TenMinuteVPN Manager is a CLI tool for managing TenMinuteVPN servers.",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
