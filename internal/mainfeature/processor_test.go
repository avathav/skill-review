package mainfeature

import (
	"reflect"
	"testing"

	"skill-review/internal/inmemmory"
)

type mockRequest struct {
	Name    string
	Message string
}

func (m mockRequest) ToRequest() Request {
	return Request{
		Name:    m.Name,
		Message: m.Message,
	}
}

func TestProcessor_Execute(t *testing.T) {
	tests := []struct {
		name    string
		args    RequestConvertable
		want    *Response
		wantErr bool
	}{
		{
			name: "test service handler response",
			args: mockRequest{
				Name:    "Test Name",
				Message: "Test Message",
			},
			want: &Response{
				Env:     "test",
				Version: "1",
				Message: "Test Message",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Processor{
				ConfigLoader: inmemmory.ConfigLoaderMock,
			}
			got, err := p.Execute(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Env, tt.want.Env) {
				t.Errorf("CheckService() got = %v, want %v", got.Env, tt.want.Env)
			}

			if !reflect.DeepEqual(got.Version, tt.want.Version) {
				t.Errorf("CheckService() got = %v, want %v", got.Version, tt.want.Version)
			}

			if !reflect.DeepEqual(got.Message, tt.want.Message) {
				t.Errorf("CheckService() got = %v, want %v", got.Message, tt.want.Message)
			}
		})
	}
}
