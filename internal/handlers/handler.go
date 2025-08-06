package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *TodoHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:id", h.Get)

	router.POST("", h.Create)

	return router
}
