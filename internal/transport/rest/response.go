package rest

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"error"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, from, problem, message string) {
	log.WithFields(log.Fields{
		"from":    from,
		"problem": problem,
	}).Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
