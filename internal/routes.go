package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/handler"
	"github.com/miyantara7/desent-solutions/internal/middleware"
)

type RouterParams struct {
	BookHandler *handler.BookHandler
}

func SetupRouter(p RouterParams) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)
	r.POST("/echo", handler.Echo)

	r.POST("/books", p.BookHandler.CreateBook)
	r.GET("/books", p.BookHandler.GetBooks)
	r.GET("/books/:id", p.BookHandler.GetBookByID)
	r.PUT("/books/:id", p.BookHandler.UpdateBook)
	r.DELETE("/books/:id", p.BookHandler.DeleteBook)

	r.POST("/auth/token", handler.GenerateToken)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/books", p.BookHandler.GetBooks)

	return r
}
