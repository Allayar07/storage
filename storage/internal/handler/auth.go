package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/model"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.Create(input)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type Request struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input Request

	if err := c.BindJSON(&input); err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.UserName, input.Password)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}