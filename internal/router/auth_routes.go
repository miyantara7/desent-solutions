package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/handler"
	"github.com/miyantara7/desent-solutions/internal/middleware"
)

type AuthRoute struct {
	BookHandler *handler.BookHandler
}

func NewAuthRoute(bh *handler.BookHandler) *AuthRoute {
	return &AuthRoute{BookHandler: bh}
}

func (ar *AuthRoute) Register(r *gin.Engine) {
	r.POST("/auth/token", handler.GenerateToken)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/books", ar.BookHandler.GetBooks)
	}
}
