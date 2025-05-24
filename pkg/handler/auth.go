package handler

import (
	"net/http"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)



func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	userId, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	wrapOkJSON(c,map[string]interface{}{
		"userId": userId,
	})
}



type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Username,input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	wrapOkJSON(c,map[string]interface{}{
		"token": token,
	})
}