package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nigelmes/todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signup)
		auth.POST("/sign-in", h.signin)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.creatList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListbyId)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.creatItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItembyId)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
