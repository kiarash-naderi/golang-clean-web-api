package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{ 

}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working!")
	return
}
func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working Post!")
	return
}
func (h *HealthHandler) HealthPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working Post by id: %s!", id))
	return
}
