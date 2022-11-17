package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) AccessPage(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		ErrorMessage(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		ErrorMessage(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		ErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("user_id", userId)
}
