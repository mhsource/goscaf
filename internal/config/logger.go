package config

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		log.WithFields(log.Fields{
			"method":      reqMethod,
			"request_uri": reqUri,
		}).Info("HTTP request log")
	}
}
