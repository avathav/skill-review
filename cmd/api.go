package cmd

import (
	"skill-review/internal/api"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "start-service",
	Short: "cli that starts reports service",
	Run:   startServices,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func startServices(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()

	api.StartAPIServer(ctx)
}
