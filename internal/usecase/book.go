package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(req *request.BookCreate) (*entity.Book, error)
	List(req *request.BookList) ([]entity.Book, *response.Meta, error)
	Get(req *request.BookGet) (*entity.Book, error)
	Update(req *request.BookUpdate) (*entity.Book, error)
	Delete(req *request.BookDelete) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *request.BookCreate) (*entity.Book, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	res, err := c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *bookUsecase) List(req *request.BookList) ([]entity.Book, *response.Meta, error) {
	res, pagination, err := c.repository.List(req.Limit, req.Page)
	if err != nil {
		return nil, nil, err
	}

	var meta = &response.Meta{
		Page:       pagination.Page,
		Total:      int(pagination.TotalRecords),
		TotalPages: pagination.TotalPage,
		PerPage:    pagination.Limit,
	}

	return res, meta, nil
}

func (c *bookUsecase) Get(req *request.BookGet) (*entity.Book, error) {
	res, err := c.repository.Find(req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *bookUsecase) Update(req *request.BookUpdate) (*entity.Book, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	res, err := c.repository.Update(req.Id, payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *bookUsecase) Delete(req *request.BookDelete) error {
	err := c.repository.Delete(req.Id)
	if err != nil {
		return err
	}

	return nil
}
