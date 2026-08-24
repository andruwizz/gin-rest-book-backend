package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(payload *dto.BookRequest) (*dto.BookResponse, error)
	Get(id string) (*dto.BookResponse, error)
	List(limit int, page int) ([]*entity.Book, *helper.Meta, error)
	Update(id string, payload *dto.BookRequest) (*dto.BookResponse, error)
	Delete(id string) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *dto.BookRequest) (*dto.BookResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	data, err := c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	response := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response, nil
}

func (c *bookUsecase) Get(id string) (*dto.BookResponse, error) {
	data, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}

	response := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response, nil
}

func (c *bookUsecase) List(limit int, page int) ([]*entity.Book, *helper.Meta, error) {
	data, pagination, err := c.repository.List(limit, page)
	if err != nil {
		return nil, nil, err
	}

	meta := &helper.Meta{
		Page:       pagination.Page,
		PerPage:    pagination.Limit,
		Total:      int(pagination.TotalRecords),
		TotalPages: pagination.TotalPage,
	}

	return data, meta, nil
}

func (c *bookUsecase) Update(id string, req *dto.BookRequest) (*dto.BookResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	data, err := c.repository.Update(id, payload)
	if err != nil {
		return nil, err
	}

	response := &dto.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response, nil
}

func (c *bookUsecase) Delete(id string) error {
	err := c.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
