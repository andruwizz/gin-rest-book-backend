package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(payload *dto.BookRequest) (*dto.BookCreateResponse, error)
	List(limit int, page int) (*dto.BookListResponse, error)
	Get(id string) (*dto.BookGetResponse, error)
	Update(id string, payload *dto.BookRequest) (*dto.BookUpdateResponse, error)
	Delete(id string) (*dto.BookDeleteResponse, error)
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *dto.BookRequest) (*dto.BookCreateResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	Data, err := c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	response := dto.BookCreateResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &response, nil
}

func (c *bookUsecase) List(limit int, page int) (*dto.BookListResponse, error) {
	Data, pagination, err := c.repository.List(limit, page)
	if err != nil {
		return nil, err
	}

	var listing []*dto.BookResponse
	for _, value := range Data {
		book := &dto.BookResponse{
			Id:        value.Id,
			Title:     value.Title,
			Author:    value.Author,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		listing = append(listing, book)
	}

	response := dto.BookListResponse{
		Success: true,
		Data:    listing,
		Meta: &dto.BookMeta{
			Page:       pagination.Page,
			PerPage:    pagination.Limit,
			Total:      int(pagination.TotalRecords),
			TotalPages: pagination.TotalPage,
		},
	}

	return &response, nil
}

func (c *bookUsecase) Get(id string) (*dto.BookGetResponse, error) {
	Data, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}

	response := dto.BookGetResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &response, nil
}

func (c *bookUsecase) Update(id string, req *dto.BookRequest) (*dto.BookUpdateResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	Data, err := c.repository.Update(id, payload)
	if err != nil {
		return nil, err
	}

	response := dto.BookUpdateResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &response, nil
}

func (c *bookUsecase) Delete(id string) (*dto.BookDeleteResponse, error) {
	err := c.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	response := dto.BookDeleteResponse{
		Success: true,
	}

	return &response, nil
}
