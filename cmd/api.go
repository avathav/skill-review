package cmd

import (
	"log"

	"skill-review/di"
	"skill-review/internal/api"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "start-api-service",
	Short: "cli that starts reports service",
	Run:   startApiService,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startApiService(_ *cobra.Command, _ []string) {
	if err := api.StartAPIServer(di.ApiPostRoutes(di.BaseParametersLoader())); err != nil {
		log.Fatalf("could not start tooling HTTP server: %v", err)
	}

}
