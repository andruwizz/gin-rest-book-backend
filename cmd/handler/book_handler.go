package handler

import (
	"net/http"
	"strconv"

	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/andruwizz/gin-book-sharing-backend/internal/model"
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
	var book model.BookRequest
	if err := ctx.ShouldBind(&book); err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	response, err := b.usecase.Create(&book)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, response)
}

func (b *BookHandler) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	response, pagination, err := b.usecase.List(limit, page)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	var meta = helper.Meta{
		Page:       pagination.Page,
		PerPage:    pagination.Limit,
		Total:      int(pagination.TotalRecords),
		TotalPages: pagination.TotalPage,
	}

	helper.OkWithMeta(ctx, response, meta)
}

func (b *BookHandler) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := b.usecase.Get(id)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, response)
}

func (b *BookHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var book model.BookRequest
	if err := ctx.ShouldBind(&book); err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	response, err := b.usecase.Update(id, &book)
	if err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	helper.OK(ctx, response)
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
