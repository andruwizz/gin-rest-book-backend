package handler

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

type InventoryHandler struct{}

func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{}
}

func (b *InventoryHandler) List(ctx *gin.Context) {
	helper.OK(ctx, gin.H{"path": "Inventory"})
}
