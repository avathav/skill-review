package cmd

import (
	"log"
	"sync"

	"skill-review/di"
	"skill-review/internal/api"
	"skill-review/internal/grpc"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "start-service",
	Short: "cli that starts reports service",
	Run:   startService,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func startService(_ *cobra.Command, _ []string) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		if err := api.StartAPIServer(di.ApiPostRoutes(di.BaseParametersLoader())); err != nil {
			log.Fatalf("could not start tooling HTTP server: %v", err)
		}

		wg.Done()
	}()

	go func() {
		if err := grpc.NewGrpcService(di.MainFeatureProcessor(di.BaseParametersLoader())); err != nil {
			log.Fatalf("could not run grpc server: %s", err.Error())
		}

		wg.Done()
	}()

	wg.Wait()

}
