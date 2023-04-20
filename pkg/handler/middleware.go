package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerpart := strings.Split(header, " ")
	if len(headerpart) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerpart[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
}
