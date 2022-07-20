package grpc

import (
	"fmt"
	"log"
	"net"

	"skill-review/internal/config"
	"skill-review/internal/grpc/services"
	"skill-review/proto"

	basegrpc "google.golang.org/grpc"
)

func NewGrpcService(c config.Loader) {
	listener, err := NewNetworkListener("tcp", "localhost:30000")
	if err != nil {
		fmt.Println("service: could not start listener:", err)
	}

	server := basegrpc.NewServer()
	checkService := services.NewCheckServiceHandler(c)
	proto.RegisterSkillReviewServer(server, *checkService)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("could not serve api")
	}
}

func NewNetworkListener(network, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Panic(err)

		return nil, fmt.Errorf("grpc: problem with setting up network listener: %s", err)
	}

	return listener, nil
}
