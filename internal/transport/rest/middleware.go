package rest

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"url":    c.Request.URL.String(),
		}).Info("get request")

		t := time.Now()

		c.Next()

		log.WithFields(log.Fields{
			"status":  c.Writer.Status(),
			"elapsed": time.Since(t).String(),
		}).Info("send response")
	}
}
