package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(req *dto.BookCreateRequest) (*dto.BookResponse, error)
	List(req *dto.BookListRequest) ([]*dto.BookResponse, *response.Meta, error)
	Get(req *dto.BookGetRequest) (*dto.BookResponse, error)
	Update(req *dto.BookUpdateRequest) (*dto.BookResponse, error)
	Delete(req *dto.BookDeleteRequest) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *dto.BookCreateRequest) (*dto.BookResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	data, err := c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	res := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return res, nil
}

func (c *bookUsecase) List(req *dto.BookListRequest) ([]*dto.BookResponse, *response.Meta, error) {
	data, pagination, err := c.repository.List(req.Limit, req.Page)
	if err != nil {
		return nil, nil, err
	}

	var res []*dto.BookResponse
	for _, value := range data {
		book := &dto.BookResponse{
			Id:        value.Id,
			Title:     value.Title,
			Author:    value.Author,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		res = append(res, book)
	}

	var meta = &response.Meta{
		Page:       pagination.Page,
		Total:      int(pagination.TotalRecords),
		TotalPages: pagination.TotalPage,
		PerPage:    pagination.Limit,
	}

	return res, meta, nil
}

func (c *bookUsecase) Get(req *dto.BookGetRequest) (*dto.BookResponse, error) {
	data, err := c.repository.Find(req.Id)
	if err != nil {
		return nil, err
	}

	res := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return res, nil
}

func (c *bookUsecase) Update(req *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	data, err := c.repository.Update(req.Id, payload)
	if err != nil {
		return nil, err
	}

	res := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return res, nil
}

func (c *bookUsecase) Delete(req *dto.BookDeleteRequest) error {
	err := c.repository.Delete(req.Id)
	if err != nil {
		return err
	}

	return nil
}
