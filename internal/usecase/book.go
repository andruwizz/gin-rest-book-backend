package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/model"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type BookUsecase interface {
	Create(payload *model.BookRequest) (*model.BookResponse, error)
	Get(id string) (*model.BookResponse, error)
	List(limit int, page int) ([]*model.BookResponse, *model.Pagination, error)
	Update(id string, payload *model.BookRequest) (*model.BookResponse, error)
	Delete(id string) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(payload *model.BookRequest) (*model.BookResponse, error) {
	data := model.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := c.repository.Create(&data)
	if err != nil {
		return nil, err
	}

	result := &model.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	return result, nil
}

func (c *bookUsecase) Get(id string) (*model.BookResponse, error) {
	book, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}

	result := &model.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	return result, nil
}

func (c *bookUsecase) List(limit int, page int) ([]*model.BookResponse, *model.Pagination, error) {
	books, pagination, err := c.repository.List(limit, page)
	if err != nil {
		return nil, nil, err
	}

	result := []*model.BookResponse{}
	for _, value := range *books {
		book := model.BookResponse{
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

func (c *bookUsecase) Update(id string, payload *model.BookRequest) (*model.BookResponse, error) {
	data := model.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := c.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	result := &model.BookResponse{
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
