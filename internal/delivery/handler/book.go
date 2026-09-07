package handler

import (
	"net/http"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/book"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase book.BookUsecase
}

func NewBookHandler(usecase book.BookUsecase) *BookHandler {
	return &BookHandler{usecase}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Input new book with title and author name
// @Tags Book
// @Accept json
// @Produce json
// @Param Body body request.BookCreate true "Request body"
// @Success 201 {object} response.DataResponse[entity.Book]
// @Failure 400 {object} response.ErrorResponse
// @Router /books [POST]
func (b *BookHandler) Create(ctx *gin.Context) {
	req := new(request.BookCreate)
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &book.BookCreateParam{
		Title:  req.Title,
		Author: req.Author,
	}
	res, err := b.usecase.Create(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusCreated, res)
}

// ListBook godoc
// @Summary List all book
// @Description Get list of books data
// @Tags Book
// @Accept json
// @Produce json
// @Param limit query int false "Item count per page"
// @Param page query int false "List page number"
// @Success 200 {object} response.ListResponse[entity.Book]
// @Failure 400 {object} response.ErrorResponse
// @Router /books [GET]
func (b *BookHandler) List(ctx *gin.Context) {
	req := new(request.BookList)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &book.BookListParam{
		Limit: req.Limit,
		Page:  req.Page,
	}

	res, pagination, err := b.usecase.List(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Collection(ctx, http.StatusOK, res, pagination)
}

// FindBook godoc
// @Summary Find a book
// @Description Get single book data
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Success 200 {object} response.DataResponse[entity.Book]
// @Failure 400 {object} response.ErrorResponse
// @Router /books/{id} [GET]
func (b *BookHandler) Find(ctx *gin.Context) {
	req := new(request.BookGet)
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &book.BookGetParam{
		Id: req.Id,
	}

	res, err := b.usecase.Get(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusOK, res)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update existing book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Param Body body request.BookUpdate true "Request body"
// @Success 200 {object} response.DataResponse[entity.Book]
// @Failure 400 {object} response.ErrorResponse
// @Router /books/{id} [PUT]
func (b *BookHandler) Update(ctx *gin.Context) {
	req := new(request.BookUpdate)
	req.Id = ctx.Param("id")

	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &book.BookUpdateParam{
		Id:     req.Id,
		Title:  req.Title,
		Author: req.Author,
	}

	res, err := b.usecase.Update(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusOK, res)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete single book data
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book Id"
// @Success 200 {object} response.EmptyResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /books/{id} [DELETE]
func (b *BookHandler) Delete(ctx *gin.Context) {
	req := new(request.BookDelete)
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &book.BookDeleteParam{
		Id: req.Id,
	}

	err := b.usecase.Delete(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Empty(ctx, http.StatusOK)
}
