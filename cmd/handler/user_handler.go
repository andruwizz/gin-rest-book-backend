package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (b *UserHandler) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"path": "User"})
}
