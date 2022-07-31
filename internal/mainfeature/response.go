package mainfeature

import (
	"time"
)

type Response struct {
	Env       string    `json:"env"`
	Message   string    `json:"param"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Echo      string    `json:"echo"`
}
