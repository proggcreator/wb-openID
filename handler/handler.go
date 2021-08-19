package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) InitRoutes() *gin.Engine {
	//router := gin.New()
	router := gin.Default()

	//midleware
	//aut := router.Group("/")

	{
		router.POST("/api/v1/all", h.showAll)
	}
	return router
}
