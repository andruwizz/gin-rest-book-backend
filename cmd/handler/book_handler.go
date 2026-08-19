package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (b *BookHandler) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"path": "Book"})
}
