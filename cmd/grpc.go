package cmd

import (
	"log"

	"skill-review/di"
	"skill-review/internal/grpc"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "start-grpc-service",
	Short: "cli that starts reports service",
	Run:   startGrpcServices,
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

func startGrpcServices(_ *cobra.Command, _ []string) {
	if err := grpc.NewGrpcService(di.MainFeatureProcessor(di.BaseParametersLoader())); err != nil {
		log.Fatalf("could not run grpc server: %s", err.Error())
	}
}
