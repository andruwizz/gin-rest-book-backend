package handler

import (
	"net/http"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase usecase.BookUsecase
}

func NewBookHandler(usecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase}
}

// Create Book godoc
// @Summary Create a new book
// @Description Input new book with title and author name
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} dto.BookResponse
// @Security ApiKeyAuth
// @Router /books [POST]
func (b *BookHandler) Create(ctx *gin.Context) {
	var req dto.BookCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Create(&req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusCreated, res, nil)
}

func (b *BookHandler) List(ctx *gin.Context) {
	var req dto.BookListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, meta, err := b.usecase.List(&req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, *meta)
}

func (b *BookHandler) Find(ctx *gin.Context) {
	var req dto.BookGetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Get(&req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, nil)
}

func (b *BookHandler) Update(ctx *gin.Context) {
	var req dto.BookUpdateRequest
	req.Id = ctx.Param("id")

	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Update(&req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, nil)
}

func (b *BookHandler) Delete(ctx *gin.Context) {
	var req dto.BookDeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	err := b.usecase.Delete(&req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, nil, nil)
}
