package service

import (
	"github.com/andruwizz/gin-collective-library-backend/internal/model"
	"github.com/andruwizz/gin-collective-library-backend/internal/repository"
)

type BookService interface {
	CreateBook(payload *model.BookRequest) (*model.BookResponse, error)
	GetBook(id string) (*model.BookResponse, error)
	ListBook(limit int, page int) ([]*model.BookResponse, *model.Pagination, error)
	UpdateBook(id string, payload *model.BookRequest) (*model.BookResponse, error)
	DeleteBook(id string) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{repository}
}

func (b *bookService) CreateBook(payload *model.BookRequest) (*model.BookResponse, error) {
	data := model.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := b.repository.Create(&data)
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

func (b *bookService) GetBook(id string) (*model.BookResponse, error) {
	book, err := b.repository.Find(id)
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

func (b *bookService) ListBook(limit int, page int) ([]*model.BookResponse, *model.Pagination, error) {
	books, pagination, err := b.repository.List(limit, page)
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

func (b *bookService) UpdateBook(id string, payload *model.BookRequest) (*model.BookResponse, error) {
	data := model.Book{
		Title:  payload.Title,
		Author: payload.Author,
	}

	book, err := b.repository.Update(id, data)
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

func (b *bookService) DeleteBook(id string) error {
	err := b.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
