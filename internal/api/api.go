package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"skill-review/internal/config"

	"github.com/gin-gonic/gin"
)

// StartAPIServer TODO move server params to config
func StartAPIServer(_ context.Context, configLoader config.Loader) {
	address := fmt.Sprintf(":%v", "40000")
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	r.POST("/:named_param", handlePostCall(configLoader))

	if runErr := r.Run(address); runErr != nil {
		log.Panicf("could not start tooling HTTP server: %v", runErr)
	}
}

func handlePostCall(configLoader config.Loader) func(gc *gin.Context) {
	return func(gc *gin.Context) {
		np := gc.Param("named_param")
		c, err := configLoader()
		if err != nil {
			gc.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not get required parameters",
			})

			return
		}

		gc.JSON(http.StatusOK, gin.H{
			"param":     np,
			"timestamp": time.Now(),
			"env":       c.Environment,
			"version":   c.Version,
		})
	}
}
