package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) creatList(c *gin.Context) {
	id, _ := c.Get("userId")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getListbyId(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
