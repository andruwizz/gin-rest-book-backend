package handler

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (b *UserHandler) List(ctx *gin.Context) {
	helper.OK(ctx, gin.H{"path": "User"})
}
