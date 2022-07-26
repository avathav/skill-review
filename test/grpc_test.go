package test

import (
	"context"
	"skill-review/di"
	"testing"
	"time"

	grpcservice "skill-review/internal/grpc"
	"skill-review/internal/inmemmory"
	proto "skill-review/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcService(t *testing.T) {
	setupGrpcServer(t)

	testCheckoutService(t)
}

func testCheckoutService(t *testing.T) {
	conn, clientErr := grpc.Dial(grpcservice.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if clientErr != nil {
		t.Fatalf("cannot create client connection: %e", clientErr)
	}

	defer conn.Close()

	client := proto.NewSkillReviewClient(conn)

	request := &proto.Request{
		Name:    "Test Name",
		Message: "Test Message",
	}

	want := &proto.Response{
		Env:     "test",
		Version: "1",
		Message: "Test Message",
	}

	response, err := client.ProcessMessage(context.Background(), request)
	if err != nil {
		t.Fatalf("check service returned error: %s", err.Error())
	}

	if response.Env != want.Env {
		t.Errorf("CheckService() got = %v, want %v", response.Env, want.Env)
	}

	if response.Version != want.Version {
		t.Errorf("CheckService() got = %v, want %v", response.Version, want.Version)
	}
}

func setupGrpcServer(t *testing.T) {
	go func() {
		if err := grpcservice.NewGrpcService(di.MainFeatureProcessor(inmemmory.ConfigLoaderMock)); err != nil {
			t.Fatalf("cannot create gerp server: %e", err)
		}
	}()

	time.Sleep(500 * time.Millisecond)
}
