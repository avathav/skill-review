package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"skill-review/di"
	"skill-review/internal/api"
	apihandlers "skill-review/internal/api/handlers"
	"skill-review/internal/inmemmory"
	"skill-review/internal/mainfeature"
)

func TestApiService(t *testing.T) {
	setupApiServer(t)

	testNamedParamService(t)
	testPostMessageService(t)
}

func testNamedParamService(t *testing.T) {
	param := "test123"
	requestURL := fmt.Sprintf("http://localhost:%d%s/%s", api.Port, api.NamedParamAddressName, param)

	response, err := http.Post(requestURL, "application/json", nil)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if response.StatusCode != 200 {
		t.Errorf("invalid status code  = %d", response.StatusCode)
	}

	byteBody, err := ioutil.ReadAll(response.Body)

	responseData := mainfeature.Response{}
	if umErr := json.Unmarshal(byteBody, &responseData); umErr != nil {
		t.Fatalf("error: %s", umErr.Error())
	}

	if responseData.Message != param {
		t.Errorf("invalid response want=%s; got=%s", param, responseData.Message)
	}

	if responseData.Env != "test" {
		t.Errorf("invalid response want=%s; got=%s", "test", responseData.Env)
	}

	if responseData.Version != "1" {
		t.Errorf("invalid response want=%s; got=%s", "1", responseData.Version)
	}

}

func testPostMessageService(t *testing.T) {
	requestURL := fmt.Sprintf("http://localhost:%d%s", api.Port, api.PostMessageAddressName)

	request := apihandlers.DefaultRequest{Message: "test345"}

	jsonRequest, mErr := json.Marshal(request)
	if mErr != nil {
		t.Fatalf("cannot marhsal request: %s", mErr.Error())
	}

	requestBody := bytes.NewReader(jsonRequest)

	response, err := http.Post(requestURL, "application/json", requestBody)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}

	if response.StatusCode != 200 {
		t.Errorf("invalid status code  = %d", response.StatusCode)
	}

	byteBody, err := ioutil.ReadAll(response.Body)

	responseData := apihandlers.DefaultResponse{}
	if umErr := json.Unmarshal(byteBody, &responseData); umErr != nil {
		t.Fatalf("error: %s", umErr.Error())
	}

	if responseData.Param != request.Message {
		t.Errorf("invalid response want=%s; got=%s", request.Message, responseData.Param)
	}

	if responseData.Env != "test" {
		t.Errorf("invalid response want=%s; got=%s", "test", responseData.Env)
	}

	if responseData.Version != "1" {
		t.Errorf("invalid response want=%s; got=%s", "1", responseData.Version)
	}

}

func setupApiServer(t *testing.T) {
	go func() {
		if err := api.StartAPIServer(di.ApiPostRoutes(inmemmory.ConfigLoaderMock)); err != nil {
			t.Fatalf("cannot create gerp server: %e", err)
		}
	}()
}
