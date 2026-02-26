package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/shared/model"
	"github.com/miyantara7/desent-solutions/internal/shared/repository/usecase"
)

type BookHandler struct {
	uc usecase.BookUsecase
}

func NewBookHandler(uc usecase.BookUsecase) *BookHandler {
	return &BookHandler{uc: uc}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req model.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	book, err := h.uc.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	author := strings.ToLower(c.Query("author"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	data := h.uc.GetBooks(author)

	start := (page - 1) * limit
	end := start + limit
	if start > len(data) {
		start = len(data)
	}
	if end > len(data) {
		end = len(data)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  data[start:end],
		"total": len(data),
		"page":  page,
		"limit": limit,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.uc.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req model.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	book, err := h.uc.UpdateBook(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
