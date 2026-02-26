package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/handler"
)

type BookRoute struct {
	Handler *handler.BookHandler
}

func NewBookRoute(h *handler.BookHandler) *BookRoute {
	return &BookRoute{Handler: h}
}

func (br *BookRoute) Register(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.POST("", br.Handler.CreateBook)
		books.GET("", br.Handler.GetBooks)
		books.GET("/:id", br.Handler.GetBookByID)
		books.PUT("/:id", br.Handler.UpdateBook)
		books.DELETE("/:id", br.Handler.DeleteBook)
	}
}
