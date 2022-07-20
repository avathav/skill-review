package services

import (
	"golang.org/x/net/context"
	"reflect"
	"skill-review/internal/config"
	"skill-review/proto"
	"testing"
)

func ConfigLoaderMock() (c config.Config, err error) {
	return config.Config{
		Environment: "test",
		Version:     "1",
	}, nil
}

func TestCheckServiceHandler_CheckService(t *testing.T) {
	tests := []struct {
		name    string
		request *proto.Request
		want    *proto.Response
		wantErr bool
	}{
		{
			name: "test service handler response",
			request: &proto.Request{
				Name:    "Test Name",
				Message: "Test Message",
			},
			want: &proto.Response{
				Env:     "test",
				Version: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CheckServiceHandler{
				ConfigLoader: ConfigLoaderMock,
			}
			got, err := s.CheckService(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Env, tt.want.Env) {
				t.Errorf("CheckService() got = %v, want %v", got.Env, tt.want.Env)
			}

			if !reflect.DeepEqual(got.Version, tt.want.Version) {
				t.Errorf("CheckService() got = %v, want %v", got.Version, tt.want.Version)
			}
		})
	}
}
