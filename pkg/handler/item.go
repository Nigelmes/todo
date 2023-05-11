package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nigelmes/todo"
	"net/http"
	"strconv"
)

func (h *Handler) creatItem(c *gin.Context) {
	Userid, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(Userid, listId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

type getAllItemResponse struct {
	Data []todo.TodoItem `json:"data"`
}

func (h *Handler) getAllItem(c *gin.Context) {
	Userid, err := getUserId(c)
	if err != nil {
		return
	}
	ListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	items, err := h.services.TodoItem.GetAll(Userid, ListId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllItemResponse{
		Data: items,
	})
}

func (h *Handler) getItembyId(c *gin.Context) {
	Userid, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	item, err := h.services.TodoItem.GetById(Userid, itemId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
