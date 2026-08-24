package handler

import (
	"net/http"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
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
	var req dto.BookCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "INVALID_DATA", "message": err.Error()},
		})
		return
	}

	res, err := b.usecase.Create(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "BAD_REQUEST", "message": err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (b *BookHandler) List(ctx *gin.Context) {
	var req dto.BookListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "INVALID_DATA", "message": err.Error()},
		})
		return
	}

	res, err := b.usecase.List(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "BAD_REQUEST", "message": err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (b *BookHandler) Find(ctx *gin.Context) {
	var req dto.BookGetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "INVALID_ID", "message": err.Error()},
		})
		return
	}

	res, err := b.usecase.Get(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "BAD_REQUEST", "message": err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (b *BookHandler) Update(ctx *gin.Context) {
	var req dto.BookUpdateRequest
	req.Id = ctx.Param("id")

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "INVALID_DATA", "message": err.Error()},
		})
		return
	}

	res, err := b.usecase.Update(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "BAD_REQUEST", "message": err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (b *BookHandler) Delete(ctx *gin.Context) {
	var req dto.BookDeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "INVALID_DATA", "message": err.Error()},
		})
		return
	}

	res, err := b.usecase.Delete(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   gin.H{"code": "BAD_REQUEST", "message": err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
