package handler

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

type LendingHandler struct{}

func NewLendingHandler() *LendingHandler {
	return &LendingHandler{}
}

func (b *LendingHandler) List(ctx *gin.Context) {
	helper.OK(ctx, gin.H{"path": "Lending"})
}
