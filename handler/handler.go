package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//midleware
	aut := router.Group("/api", h.userIdentity)

	{
		aut.POST("/api/v1/all", h.showAll)
	}
	return router
}
