package handlers

import (
	"net/http"

	"skill-review/internal/mainfeature"

	"github.com/gin-gonic/gin"
)

type NamedParam struct {
	processor *mainfeature.Processor
}

func (h NamedParam) Handler(gc *gin.Context) {
	defaultRequest := DefaultRequest{Message: gc.Param("message")}

	response, err := h.processor.Execute(defaultRequest)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not process request",
		})
	}

	gc.JSON(http.StatusOK, convertToDefaultResponse(*response))
}

func NewNamedParamHandler(processor *mainfeature.Processor) *NamedParam {
	return &NamedParam{processor: processor}
}
