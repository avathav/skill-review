package grpc

import (
	"fmt"
	"net"

	"skill-review/internal/mainfeature"
	proto "skill-review/proto"

	"github.com/pkg/errors"
	basegrpc "google.golang.org/grpc"
)

const (
	ServerAddress = "localhost:40000"
)

func NewGrpcService(p *mainfeature.Processor) error {
	listener, err := NewNetworkListener("tcp", ServerAddress)
	if err != nil {
		return errors.Wrap(err, "service: could not start listener")
	}

	server := basegrpc.NewServer()
	proto.RegisterSkillReviewServer(server, NewSkillReviewServer(p))

	if err = server.Serve(listener); err != nil {
		return errors.Wrap(err, "could not serve api")
	}

	return nil
}

func NewNetworkListener(network, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, fmt.Errorf("grpc: problem with setting up network listener: %s", err)
	}

	return listener, nil
}
