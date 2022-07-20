package services

import (
	"skill-review/internal/config"
	"time"

	"skill-review/proto"

	"golang.org/x/net/context"
)

type Handling interface {
	CheckService(context.Context, *proto.Request) (*proto.Response, error)
}

type CheckServiceHandler struct {
	ConfigLoader config.Loader
}

func NewCheckServiceHandler(l config.Loader) *CheckServiceHandler {
	return &CheckServiceHandler{ConfigLoader: l}
}

func (s CheckServiceHandler) CheckService(_ context.Context, request *proto.Request) (*proto.Response, error) {
	c, err := s.ConfigLoader()
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		Env:       c.Environment,
		Timestamp: time.Now().String(),
		Version:   c.Version,
	}, nil
}
