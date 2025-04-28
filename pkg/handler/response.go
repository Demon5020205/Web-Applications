package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

type statusResponce struct {
	Status string `json:"status"`
}

func newErrorResponce(c *gin.Context, statusCode int, messege string) {
	logrus.Error(messege)
	c.AbortWithStatusJSON(statusCode, errorResponce{messege})
}