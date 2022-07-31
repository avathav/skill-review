package grpc

import (
	"skill-review/internal/mainfeature"
	proto "skill-review/proto"

	"golang.org/x/net/context"
)

type Handling interface {
	CheckService(context.Context, *proto.Request) (*proto.Response, error)
}

type CheckServiceRequest struct {
	Name    string
	Message string
}

func (c CheckServiceRequest) ToRequest() mainfeature.Request {
	return mainfeature.Request{
		Name:    c.Name,
		Message: c.Message,
	}
}

type CheckServiceHandler struct {
	processor *mainfeature.Processor
}

func NewSkillReviewServer(p *mainfeature.Processor) *CheckServiceHandler {
	return &CheckServiceHandler{processor: p}
}

func (s CheckServiceHandler) ProcessMessage(_ context.Context, req *proto.Request) (*proto.Response, error) {
	res, err := s.processor.Execute(CheckServiceRequest{
		Name:    req.Name,
		Message: req.Message,
	})
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		Env:       res.Env,
		Timestamp: res.Timestamp.String(),
		Version:   res.Version,
		Message:   res.Message,
		Echo:      res.Echo,
	}, nil
}
