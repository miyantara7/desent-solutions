package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/handler"
)

type SpeedRunRoute struct {
	handler *handler.SpeedrunHandler
}

func NewSpeedRunRoute(handler *handler.SpeedrunHandler) *SpeedRunRoute {
	return &SpeedRunRoute{
		handler: handler,
	}
}

func (h *SpeedRunRoute) Register(r *gin.Engine) {
	r.POST("/speedrun", h.handler.Run)
}
