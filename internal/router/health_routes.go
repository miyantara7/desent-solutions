package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/handler"
)

type HealthRoute struct{}

func NewHealthRoute() *HealthRoute {
	return &HealthRoute{}
}

func (h *HealthRoute) Register(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	r.POST("/echo", handler.Echo)
}
