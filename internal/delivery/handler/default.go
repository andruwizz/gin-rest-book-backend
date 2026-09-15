package handler

import (
	"errors"
	"net/http"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/gin-gonic/gin"
)

type DefaultHandler struct{}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (d *DefaultHandler) Index(ctx *gin.Context) {
	response.Empty(ctx, http.StatusOK)
}

func (d *DefaultHandler) NotFound(ctx *gin.Context) {
	err := errors.New("not found")
	response.ErrorResource(ctx, http.StatusNotFound, err)
}
