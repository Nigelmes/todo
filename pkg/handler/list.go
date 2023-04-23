package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nigelmes/todo"
	"net/http"
)

func (h *Handler) creatList(c *gin.Context) {
	Userid, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(Userid, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
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
