package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorization = "Authorization"
	userCtx = "userId"
)


func (h* Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorization)  
	if header == "" {
		newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
		return
	}

	headerParts := strings.Split(header," ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth headeer")
	}

	userId, err :=  h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)

} 