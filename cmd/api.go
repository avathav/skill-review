package cmd

import (
	"github.com/spf13/cobra"
	"skill-review/di"
	"skill-review/internal/api"
)

var apiCmd = &cobra.Command{
	Use:   "start-api-service",
	Short: "cli that starts reports service",
	Run:   startApiService,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startApiService(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()

	api.StartAPIServer(ctx, di.BaseParametersLoader())

}
