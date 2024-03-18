package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tenminutevpn/tenminutevpn-manager/squid"
)

func squidSetup(port int) {
	s := squid.NewSquid(port)

	err := s.Write("/etc/squid/squid.conf")
	if err != nil {
		log.Fatalf("Error writing squid config: %s", err)
		return
	}

	err = s.SystemdService().Enable()
	if err != nil {
		log.Fatalf("Error enabling squid service: %s", err)
		return
	}

	err = s.SystemdService().Start()
	if err != nil {
		log.Fatalf("Error starting squid service: %s", err)
		return
	}
}

var squidSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Squid Proxy",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		squidSetup(port)
	},
}

var squidCmd = &cobra.Command{
	Use:   "squid",
	Short: "Squid Proxy Management",
}

func init() {
	squidSetupCmd.Flags().IntP("port", "p", 3128, "Port to listen on")
	squidCmd.AddCommand(squidSetupCmd)

	rootCmd.AddCommand(squidCmd)
}
