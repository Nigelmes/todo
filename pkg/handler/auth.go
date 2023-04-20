package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nigelmes/todo"
	"net/http"
)

func (h *Handler) signup(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signin(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
