package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kiarash-naderi/golang-clean-web-api/api/routers"
	"github.com/kiarash-naderi/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}