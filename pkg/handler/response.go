package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode,Error{Message: message})
}

func wrapOkJSON(c *gin.Context, response map[string]interface{}) {
	c.JSON(http.StatusOK,response)
}