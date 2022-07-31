package mainfeature

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"skill-review/internal/config"

	"github.com/pkg/errors"
)

type Processor struct {
	ConfigLoader config.Loader
}

func NewProcessor(configLoader config.Loader) *Processor {
	return &Processor{ConfigLoader: configLoader}
}

type RequestConvertable interface {
	ToRequest() Request
}

func (p Processor) Execute(requestMessage RequestConvertable) (*Response, error) {
	r := requestMessage.ToRequest()

	c, err := p.ConfigLoader()
	if err != nil {
		return nil, err
	}

	echoRes, echoErr := processPostmanEchoCall(r)
	if echoErr != nil {
		return nil, echoErr
	}

	return &Response{
		Env:       c.Environment,
		Message:   r.Message,
		Timestamp: time.Now(),
		Version:   c.Version,
		Echo:      string(echoRes),
	}, nil
}

func processPostmanEchoCall(r Request) ([]byte, error) {
	jsonRequest, mErr := json.Marshal(r)
	if mErr != nil {
		return nil, errors.Wrap(mErr, "cannot marshal request")
	}

	resp, pmErr := http.Post("https://postman-echo.com/post/", "application/json", bytes.NewReader(jsonRequest))
	if pmErr != nil {
		return nil, errors.Wrap(pmErr, "cannot ping postman-echo service")
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
