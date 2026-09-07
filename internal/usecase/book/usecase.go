package book

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
)

type BookUsecase interface {
	Create(req *BookCreateParam) (*entity.Book, error)
	List(req *BookListParam) ([]entity.Book, *entity.Pagination, error)
	Get(req *BookGetParam) (*entity.Book, error)
	Update(req *BookUpdateParam) (*entity.Book, error)
	Delete(req *BookDeleteParam) error
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return &bookUsecase{repository}
}

func (c *bookUsecase) Create(req *BookCreateParam) (*entity.Book, error) {
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

func (c *bookUsecase) List(req *BookListParam) ([]entity.Book, *entity.Pagination, error) {
	res, pagination, err := c.repository.List(req.Limit, req.Page)
	if err != nil {
		return nil, nil, err
	}

	return res, pagination, nil
}

func (c *bookUsecase) Get(req *BookGetParam) (*entity.Book, error) {
	book := new(entity.Book)
	err := c.repository.Find(req.Id, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (c *bookUsecase) Update(req *BookUpdateParam) (*entity.Book, error) {
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

func (c *bookUsecase) Delete(req *BookDeleteParam) error {
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
