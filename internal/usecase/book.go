package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/dto"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(payload *dto.BookRequest) (*dto.BookResponse, error)
	Get(id string) (*dto.BookResponse, error)
	List(limit int, page int) ([]*dto.BookResponse, *dto.Pagination, error)
	Update(id string, payload *dto.BookRequest) (*dto.BookResponse, error)
	Delete(id string) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(payload *dto.BookRequest) (*dto.BookResponse, error) {
	data := entity.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := c.repository.Create(&data)
	if err != nil {
		return nil, err
	}

	result := &dto.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	return result, nil
}

func (c *bookUsecase) Get(id string) (*dto.BookResponse, error) {
	book, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}

	result := &dto.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	return result, nil
}

func (c *bookUsecase) List(limit int, page int) ([]*dto.BookResponse, *dto.Pagination, error) {
	books, pagination, err := c.repository.List(limit, page)
	if err != nil {
		return nil, nil, err
	}

	result := []*dto.BookResponse{}
	for _, value := range *books {
		book := dto.BookResponse{
			Id:        value.Id,
			Title:     value.Title,
			Author:    value.Author,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}

		result = append(result, &book)
	}

	return result, pagination, nil
}

func (c *bookUsecase) Update(id string, payload *dto.BookRequest) (*dto.BookResponse, error) {
	data := entity.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := c.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	result := &dto.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	return result, nil
}

func (c *bookUsecase) Delete(id string) error {
	err := c.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
