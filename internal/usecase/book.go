package usecase

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
)

type BookUsecase interface {
	Create(req *request.BookCreate) (*entity.Book, error)
	List(req *request.BookList) ([]entity.Book, *entity.Pagination, error)
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
	book := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	err := c.repository.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (c *bookUsecase) List(req *request.BookList) ([]entity.Book, *entity.Pagination, error) {
	res, pagination, err := c.repository.List(req.Limit, req.Page)
	if err != nil {
		return nil, nil, err
	}

	return res, pagination, nil
}

func (c *bookUsecase) Get(req *request.BookGet) (*entity.Book, error) {
	book := new(entity.Book)
	err := c.repository.Find(req.Id, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (c *bookUsecase) Update(req *request.BookUpdate) (*entity.Book, error) {
	book := new(entity.Book)
	err := c.repository.Find(req.Id, book)
	if err != nil {
		return book, err
	}

	book.Title = req.Title
	book.Author = req.Author

	err = c.repository.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (c *bookUsecase) Delete(req *request.BookDelete) error {
	book := new(entity.Book)
	err := c.repository.Find(req.Id, book)
	if err != nil {
		return err
	}

	err = c.repository.Delete(book)
	if err != nil {
		return err
	}

	return nil
}
