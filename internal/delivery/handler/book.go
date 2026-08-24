package handler

import (
	"net/http"
	"strconv"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/andruwizz/gin-book-sharing-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase usecase.BookUsecase
}

func NewBookHandler(usecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase}
}

func (b *BookHandler) Create(ctx *gin.Context) {
	var req dto.BookRequest
	if err := ctx.ShouldBind(&req); err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Create(&req)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, res)
}

func (b *BookHandler) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res, meta, err := b.usecase.List(limit, page)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OkWithMeta(ctx, res, *meta)
}

func (b *BookHandler) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := b.usecase.Get(id)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, res)
}

func (b *BookHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.BookRequest
	if err := ctx.ShouldBind(&req); err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Update(id, &req)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, res)
}

func (b *BookHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := b.usecase.Delete(id)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, nil)
}
