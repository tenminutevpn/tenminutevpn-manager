package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func setup() {
	fmt.Println("Setting up TenMinuteVPN")
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
