package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) showAll(c *gin.Context) {

	fmt.Fprintln(c.Writer, "Helloooo")
}
