package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/andruwizz/gin-collective-library-backend/internal/helper"
	"github.com/andruwizz/gin-collective-library-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (b *BookHandler) Create(ctx *gin.Context) {
	var book model.BookRequest
	if err := ctx.ShouldBind(&book); err != nil {
		helper.Fail(ctx, http.StatusBadRequest, "INVALID_DATA", err.Error())
		return
	}

	var response = model.BookResponse{
		Id:    "BOOK01212421",
		Title: book.Title,
		Author: model.AuthorProfile{
			Id:   book.AuthorId,
			Name: "Andrea Hirata",
		},
		CreatedAt: time.RFC3339,
		UpdatedAt: time.RFC3339,
	}

	helper.OK(ctx, response)
}

func (b *BookHandler) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var response = []model.BookResponse{
		{
			Id:    "BOOK01212421",
			Title: "Book One",
			Author: model.AuthorProfile{
				Id:   "AU03232452",
				Name: "Andrea Hirata",
			},
			CreatedAt: time.RFC3339,
			UpdatedAt: time.RFC3339,
		},
		{
			Id:    "BOOK01212422",
			Title: "Book Two",
			Author: model.AuthorProfile{
				Id:   "AU03232452",
				Name: "Andrea Hirata",
			},
			CreatedAt: time.RFC3339,
			UpdatedAt: time.RFC3339,
		},
	}

	var meta = helper.Meta{
		Page:       page,
		PerPage:    10,
		Total:      2,
		TotalPages: 1,
	}

	helper.OkWithMeta(ctx, response, meta)
}

func (b *BookHandler) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	var response = model.BookResponse{
		Id:    id,
		Title: "Book One",
		Author: model.AuthorProfile{
			Id:   "AU03232452",
			Name: "Andrea Hirata",
		},
		CreatedAt: time.RFC3339,
		UpdatedAt: time.RFC3339,
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

	var response = model.BookResponse{
		Id:    id,
		Title: book.Title,
		Author: model.AuthorProfile{
			Id:   book.AuthorId,
			Name: "Andrea Hirata",
		},
		CreatedAt: time.RFC3339,
		UpdatedAt: time.RFC3339,
	}

	helper.OK(ctx, response)
}

func (b *BookHandler) Delete(ctx *gin.Context) {
	helper.OK(ctx, nil)
}
