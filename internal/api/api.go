package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Port                   = 30000
	NamedParamAddressName  = "/named_param"
	PostMessageAddressName = "/message"
)

type Route struct {
	Method  string
	Address string
	Handler gin.HandlerFunc
}

// StartAPIServer TODO move server params to config
func StartAPIServer(routes []Route) error {
	address := fmt.Sprintf(":%v", Port)
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.GET(route.Address, route.Handler)
		case "POST":
			r.POST(route.Address, route.Handler)
		}
	}

	return r.Run(address)
}
