package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"skill-review/internal/mainfeature"

	"github.com/gin-gonic/gin"
)

type ApiHandler interface {
	Handler(gc *gin.Context)
}

type DefaultRequest struct {
	Message string `json:"messsage"`
}

type DefaultResponse struct {
	Env       string    `json:"env"`
	Param     string    `json:"param"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Echo      string    `json:"echo"`
}

func (d DefaultRequest) ToRequest() mainfeature.Request {
	return mainfeature.Request{
		Message: d.Message,
	}
}

type Default struct {
	processor *mainfeature.Processor
}

//
// @Summary Provides default json message processing
// @Description reads JSON message
// @Accept  json
// @Param request body DefaultRequest true "request message"
// @Produce  json
// @Success 200 {object} DefaultResponse "ok"
// @Router /message [post]
func (d Default) Handler(gc *gin.Context) {
	body, readErr := ioutil.ReadAll(gc.Request.Body)
	if readErr != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get request",
		})
	}

	defaultRequest := DefaultRequest{}
	if umErr := json.Unmarshal(body, &defaultRequest); umErr != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get request",
		})
	}

	response, err := d.processor.Execute(defaultRequest)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not process request",
		})
	}

	gc.JSON(http.StatusOK, convertToDefaultResponse(*response))
}

func NewDefaultHandler(processor *mainfeature.Processor) *Default {
	return &Default{processor: processor}
}

func convertToDefaultResponse(mr mainfeature.Response) DefaultResponse {
	return DefaultResponse{
		Env:       mr.Env,
		Param:     mr.Message,
		Timestamp: mr.Timestamp,
		Version:   mr.Version,
		Echo:      mr.Echo,
	}
}
