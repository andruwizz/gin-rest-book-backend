package handler

import (
	"net/http"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
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

// CreateBook godoc
// @Summary Create a new book
// @Description Input new book with title and author name
// @Tags Book
// @Accept json
// @Produce json
// @Param Body body request.BookCreate true "Request body"
// @Success 201 {object} response.ApiDataResponse{data=entity.Book}
// @Failure 400 {object} response.ApiErrorResponse
// @Router /books [POST]
func (b *BookHandler) Create(ctx *gin.Context) {
	req := new(request.BookCreate)
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Create(req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusCreated, res, nil)
}

// ListBook godoc
// @Summary List all book
// @Description Get list of books data
// @Tags Book
// @Accept json
// @Produce json
// @Param limit query int false "Item count per page"
// @Param page query int false "List page number"
// @Success 200 {object} response.ApiDataResponse{data=entity.Book}
// @Failure 400 {object} response.ApiErrorResponse
// @Router /books [GET]
func (b *BookHandler) List(ctx *gin.Context) {
	req := new(request.BookList)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, meta, err := b.usecase.List(req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, *meta)
}

// FindBook godoc
// @Summary Find a book
// @Description Get single book data
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Success 200 {object} response.ApiDataResponse{data=entity.Book}
// @Failure 400 {object} response.ApiErrorResponse
// @Router /books/{id} [GET]
func (b *BookHandler) Find(ctx *gin.Context) {
	req := new(request.BookGet)
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Get(req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, nil)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update existing book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Param Body body request.BookUpdate true "Request body"
// @Success 200 {object} response.ApiDataResponse{data=entity.Book}
// @Failure 400 {object} response.ApiErrorResponse
// @Router /books/{id} [PUT]
func (b *BookHandler) Update(ctx *gin.Context) {
	req := new(request.BookUpdate)
	req.Id = ctx.Param("id")

	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	res, err := b.usecase.Update(req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, res, nil)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete single book data
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Success 200 {object} response.ApiEmptyResponse
// @Failure 400 {object} response.ApiErrorResponse
// @Router /books/{id} [DELETE]
func (b *BookHandler) Delete(ctx *gin.Context) {
	req := new(request.BookDelete)
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	err := b.usecase.Delete(req)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, nil, nil)
}
