package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(req *dto.BookCreateRequest) (*dto.BookCreateResponse, error)
	List(req *dto.BookListRequest) (*dto.BookListResponse, error)
	Get(req *dto.BookGetRequest) (*dto.BookGetResponse, error)
	Update(req *dto.BookUpdateRequest) (*dto.BookUpdateResponse, error)
	Delete(req *dto.BookDeleteRequest) (*dto.BookDeleteResponse, error)
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *dto.BookCreateRequest) (*dto.BookCreateResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	Data, err := c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	res := dto.BookCreateResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &res, nil
}

func (c *bookUsecase) List(req *dto.BookListRequest) (*dto.BookListResponse, error) {
	Data, pagination, err := c.repository.List(req.Limit, req.Page)
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

	res := dto.BookListResponse{
		Success: true,
		Data:    listing,
		Meta: &dto.BookMeta{
			Page:       pagination.Page,
			PerPage:    pagination.Limit,
			Total:      int(pagination.TotalRecords),
			TotalPages: pagination.TotalPage,
		},
	}

	return &res, nil
}

func (c *bookUsecase) Get(req *dto.BookGetRequest) (*dto.BookGetResponse, error) {
	Data, err := c.repository.Find(req.Id)
	if err != nil {
		return nil, err
	}

	res := dto.BookGetResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &res, nil
}

func (c *bookUsecase) Update(req *dto.BookUpdateRequest) (*dto.BookUpdateResponse, error) {
	payload := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	Data, err := c.repository.Update(req.Id, payload)
	if err != nil {
		return nil, err
	}

	res := dto.BookUpdateResponse{
		Success: true,
		Data: &dto.BookResponse{
			Id:        Data.Id,
			Title:     Data.Title,
			Author:    Data.Author,
			CreatedAt: Data.CreatedAt,
			UpdatedAt: Data.UpdatedAt,
		},
	}

	return &res, nil
}

func (c *bookUsecase) Delete(req *dto.BookDeleteRequest) (*dto.BookDeleteResponse, error) {
	err := c.repository.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	res := dto.BookDeleteResponse{
		Success: true,
	}

	return &res, nil
}
