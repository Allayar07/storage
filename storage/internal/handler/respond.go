package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MessageError struct {
	Message string `json:"message"`
}

func ErrorMessage(c *gin.Context, statuCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statuCode, MessageError{
		Message: message,
	})
}
