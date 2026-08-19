package handler

import (
	"github.com/andruwizz/gin-collective-library-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (b *BookHandler) List(ctx *gin.Context) {
	helper.OK(ctx, gin.H{"path": "Book"})
}
