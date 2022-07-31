package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title SkillReviewApp
// @version 1.0
// @description Rest and Grpc endpoints.
// @BasePath /
func StartAPIServer(routes []Route) error {
	address := fmt.Sprintf(":%v", Port)
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	registerRoutes(r, routes)

	return r.Run(address)
}

func registerRoutes(r *gin.Engine, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.GET(route.Address, route.Handler)
		case "POST":
			r.POST(route.Address, route.Handler)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
